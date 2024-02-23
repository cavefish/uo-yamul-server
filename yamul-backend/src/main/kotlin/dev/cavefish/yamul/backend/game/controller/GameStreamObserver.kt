package dev.cavefish.yamul.backend.game.controller


import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgApplyWorldPatches
import dev.cavefish.yamul.backend.game.api.MsgCharacterSelection
import dev.cavefish.yamul.backend.game.api.MsgHealthBar
import dev.cavefish.yamul.backend.game.api.MsgMapChange
import dev.cavefish.yamul.backend.game.api.MsgPlayMusic
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import io.grpc.stub.StreamObserver
import org.tinylog.Logger

class GameStreamObserver(
    private val outputStream: StreamObserver<StreamPackage>
) :
    StreamObserver<StreamPackage> {

    override fun onNext(message: StreamPackage?) {
        if (message == null) return
        when (message.type) {
            MsgType.TypeCharacterSelection -> process(message.body.characterSelection, this::onCharacterSelected)
            null -> throw Exception("message cannot be null")
            else -> unimplementedMessage(message)
        }
    }

    private fun <T> process(obj: T, function: (T) -> Unit) {
        Logger.debug(obj)
        function(obj)
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

    private fun onCharacterSelected(ignored: MsgCharacterSelection) {
        send(MsgType.TypeApplyWorldPatches) {it.setApplyWorldPatches(MsgApplyWorldPatches.getDefaultInstance())}
        send(MsgType.TypePlayMusic) {it.setPlayMusic(MsgPlayMusic.getDefaultInstance())}
        send(MsgType.TypeMapChange) {it.setMapChange(MsgMapChange.getDefaultInstance())}
        send(MsgType.TypeHealthBar) {it.setHealthBar(MsgHealthBar.getDefaultInstance())}

    }
    private fun send(msgType: MsgType, f: (Message.Builder) -> Unit) {
        outputStream.onNext(StreamPackage.newBuilder().setType(msgType).setBody(Message.newBuilder().apply(f)).build())
    }
}