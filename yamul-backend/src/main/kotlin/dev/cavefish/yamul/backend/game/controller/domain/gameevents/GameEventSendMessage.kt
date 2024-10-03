package dev.cavefish.yamul.backend.game.controller.domain.gameevents

import dev.cavefish.yamul.backend.game.api.MsgSystemSendText
import dev.cavefish.yamul.backend.game.api.MsgSystemSendText_Type
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamObserver
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.domain.Hues
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser

data class GameEventSendMessage(val message: String, override val filter: GameEventFilter = GameEventFilter.ANY) :
    GameEvent(filter) {
    override fun invoke(state: GameState, streamWrapper: GameStreamWrapper) {
        streamWrapper.send(MsgType.TypeSystemSendText) {
            it.setSystemSendText(
                MsgSystemSendText.newBuilder()
                    // TODO remove hardcoded values
                    .setHue(Hues.Red.hue.toUInt16().toInt())
                    .setFont(1)
                    .setType(MsgSystemSendText_Type.system)
                    .setName("SYSTEM")
                    .setBody(message)
            )
        }
    }
}