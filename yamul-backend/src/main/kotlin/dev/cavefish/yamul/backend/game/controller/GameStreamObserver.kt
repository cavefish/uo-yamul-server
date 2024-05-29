package dev.cavefish.yamul.backend.game.controller


import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.processors.MessageProcessor
import io.grpc.stub.StreamObserver
import org.tinylog.Logger

class GameStreamObserver(
    private val outputStream: StreamObserver<StreamPackage>,
    private val processors: Map<MsgType, MessageProcessor<*>>
) : StreamObserver<StreamPackage>, GameStreamWrapper {

    private var gameState: GameState = GameState()

    override fun onNext(message: StreamPackage?) {
        if (message == null) return
        val messageProcessor = processors[message.type]
        if (messageProcessor == null) unimplementedMessage(message)
        else gameState = messageProcessor.process(message, gameState,this)
    }

    override fun onError(p0: Throwable?) {
        Logger.error(p0)
    }

    override fun onCompleted() {
        Logger.info("Game stream closed", "")
    }

    private fun unimplementedMessage(message: StreamPackage) {
        TODO("Unimplemented message type %s".format(message.type.name))
    }

    override fun send(msgType: MsgType, f: (Message.Builder) -> Unit) {
        outputStream.onNext(StreamPackage.newBuilder().setType(msgType).setBody(Message.newBuilder().apply(f)).build())
    }
}