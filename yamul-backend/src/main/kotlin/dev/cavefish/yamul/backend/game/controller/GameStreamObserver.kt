package dev.cavefish.yamul.backend.game.controller


import dev.cavefish.yamul.backend.Constants
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameEvent
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventObserver
import dev.cavefish.yamul.backend.game.controller.processors.MessageProcessor
import io.grpc.StatusRuntimeException
import io.grpc.stub.StreamObserver
import org.tinylog.kotlin.Logger
import java.util.concurrent.atomic.AtomicReference

class GameStreamObserver(
    private val outputStream: StreamObserver<StreamPackage>,
    private val processors: Map<MsgType, MessageProcessor<*>>,
    private val gameStreamEventCoordinator: GameStreamEventCoordinator
) : StreamObserver<StreamPackage>, GameStreamWrapper, GameStreamEventObserver {

    private var gameStateRegister: AtomicReference<GameState> = AtomicReference()

    override fun onNext(message: StreamPackage?) {
        if (message == null) return
        val messageProcessor = processors[message.type]
        if (messageProcessor == null) {
            unimplementedMessage(message)
            return
        }
        gameStateRegister.getAndUpdate { oldState ->
            val nextState =
                messageProcessor.process(message, oldState, Constants.AUTH_CONTEXT_LOGGED_USER.get(), this)
            if (nextState != null && oldState == null) {
                gameStreamEventCoordinator.subscribe(this)
            }
            if (nextState == null) {
                gameStreamEventCoordinator.unsubscribe(this)
            }
            nextState
        }
    }

    override fun onError(errr: Throwable?) {
        when (errr) {
            is StatusRuntimeException -> Logger.warn("%s %s".format(errr.status, errr.message), errr)
            else -> {
                Logger.error(errr)
            }
        }
    }

    override fun onCompleted() {
        gameStreamEventCoordinator.unsubscribe(this)
        outputStream.onCompleted()
        Logger.info("Game stream closed")
    }

    private fun unimplementedMessage(message: StreamPackage) {
        TODO("Unimplemented message type %s".format(message.type.name))
    }

    override fun send(msgType: MsgType, f: (Message.Builder) -> Unit) {
        val body = StreamPackage.newBuilder().setType(msgType).setBody(Message.newBuilder().apply(f)).build()
        // Logger.debug(body)
        outputStream.onNext(body)
    }

    override fun onEvent(event: GameEvent) {
        if (!event.appliesTo(gameStateRegister.get())) return
        Logger.debug("Processing event: {}", event)
        event(this.gameStateRegister.get(), this)
    }
}
