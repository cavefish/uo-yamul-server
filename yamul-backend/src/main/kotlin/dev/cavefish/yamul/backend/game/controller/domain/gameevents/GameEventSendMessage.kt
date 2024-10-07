package dev.cavefish.yamul.backend.game.controller.domain.gameevents

import dev.cavefish.yamul.backend.common.api.Fonts
import dev.cavefish.yamul.backend.common.api.MessageType
import dev.cavefish.yamul.backend.game.api.MsgSystemSendText
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.UserMessageType
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.mappers.ObjectIdMapper

data class GameEventSendMessage(
    val message: String,
    val type: UserMessageType,
    val origin: GameObject?,
    override val filter: GameEventFilter = GameEventFilter.ANY
) :
    GameEvent(filter) {
    override fun invoke(state: State, streamWrapper: GameStreamWrapper) {
        val message = MsgSystemSendText.newBuilder()
            // TODO remove hardcoded values
            .setHue(0x02B2)
            .setFont(Fonts.Font_normal)
            .setBody(message)

        if (origin != null) {
            message
                .setId(ObjectIdMapper.INSTANCE.create(origin.id))
                .setType(type.apiValue)
            origin.name ?: message.setName(origin.name)
        } else {
            message.setName("SYSTEM")
                .setType(MessageType.MessageType_system)
        }
        streamWrapper.send(MsgType.TypeSystemSendText) {
            it.setSystemSendText(
                message
            )
        }
    }
}