package dev.cavefish.yamul.backend.game.controller


import dev.cavefish.yamul.backend.Constants
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameEvent
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventObserver
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateClosesConnection
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateError
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateInitial
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateLoggedOut
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateProperty
import dev.cavefish.yamul.backend.game.controller.processors.MessageProcessor
import io.grpc.StatusRuntimeException
import io.grpc.stub.StreamObserver
import kotlinx.coroutines.runBlocking
import org.tinylog.kotlin.Logger
import java.util.concurrent.atomic.AtomicReference

class GameStreamObserver(
    private val outputStream: StreamObserver<StreamPackage>,
    private val processors: Map<MsgType, MessageProcessor<*>>,
    private val gameStreamEventCoordinator: GameStreamEventCoordinator
) : StreamObserver<StreamPackage>, GameStreamWrapper, GameStreamEventObserver {

    private var stateRegister: AtomicReference<State> = AtomicReference()

    override fun onNext(message: StreamPackage?) {
        if (stateRegister.get() == null) {
            stateRegister.getAndUpdate {
                it ?: StateInitial(Constants.AUTH_CONTEXT_LOGGED_USER.get())
            }
        }
        if (message == null) return
        val messageProcessor = processors[message.type]
        if (messageProcessor == null) {
            unimplementedMessage(message)
            return
        }
        stateRegister.getAndUpdate { oldState ->
            messageProcessor.processStreamPackage(message, oldState, this).also {
                onStateChange(oldState, it)
            }
        }
    }

    private fun <T:State> onStateChange(
        from: State,
        to: T
    ):T {
        if (to.hasWonProperty(from, StateProperty.RECEIVE_EVENTS)) {
            gameStreamEventCoordinator.subscribe(this)
        }
        if (to.hasLostProperty(from, StateProperty.RECEIVE_EVENTS)) {
            gameStreamEventCoordinator.unsubscribe(this)
        }
        if (to is StateError) Logger.error(to.errorDescription)
        if (to is StateClosesConnection) {
            outputStream.onCompleted()
            Logger.info("Game stream closed")
        }
        return to
    }

    override fun onError(errr: Throwable?) {
        when (errr) {
            is StatusRuntimeException -> Logger.warn("%s %s".format(errr.status, errr.message), errr)
            else -> {
                Logger.error(errr)
            }
        }
        stateRegister.getAndUpdate {
            onStateChange(it, StateError(errr.toString()))
        }
    }

    override fun onCompleted() {
        stateRegister.getAndUpdate {
            onStateChange(it, StateLoggedOut)
        }
    }

    private fun unimplementedMessage(message: StreamPackage) {
        Logger.error("Unimplemented message type {}", message.type.name)
    }

    override fun send(msgType: MsgType, f: (Message.Builder) -> Unit) {
        val body = StreamPackage.newBuilder().setType(msgType).setBody(Message.newBuilder().apply(f)).build()
        Logger.debug(body)
        outputStream.onNext(body)
    }

    override fun onEvent(event: GameEvent) {
        val state = stateRegister.get()
        if (!state.hasProperty(StateProperty.RECEIVE_EVENTS)) return
        if (!event.appliesTo(state)) return
        Logger.debug("Processing event: {}", event)
        event(state, this)
    }
}
