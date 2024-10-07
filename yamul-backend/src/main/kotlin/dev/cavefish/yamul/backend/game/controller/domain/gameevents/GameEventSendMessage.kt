package dev.cavefish.yamul.backend.game.controller.domain.gameevents

import dev.cavefish.yamul.backend.common.api.Fonts
import dev.cavefish.yamul.backend.common.api.MessageType
import dev.cavefish.yamul.backend.game.api.MsgSystemSendText
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State

data class GameEventSendMessage(val message: String, override val filter: GameEventFilter = GameEventFilter.ANY) :
    GameEvent(filter) {
    override fun invoke(state: State, streamWrapper: GameStreamWrapper) {
        streamWrapper.send(MsgType.TypeSystemSendText) {
            it.setSystemSendText(
                MsgSystemSendText.newBuilder()
                    // TODO remove hardcoded values
                    .setHue(0x02B2)
                    .setFont(Fonts.Font_gothic)
                    .setType(MessageType.MessageType_system)
                    .setName("SYSTEM")
                    .setBody(message)
            )
        }
    }
}