package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.MsgUnicodeSpeechSelected
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameEventSendMessage
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import org.springframework.stereotype.Component

@Component
class OnUnicodeSpeechSelected (
    private val eventCoordinator: GameStreamEventCoordinator
) : MessageProcessor<MsgUnicodeSpeechSelected>(
    MsgType.TypeUnicodeSpeechSelected, Message::getUnicodeSpeechSelected
) {

    override fun process(payload: MsgUnicodeSpeechSelected, state: State, wrapper: GameStreamWrapper): State {
        eventCoordinator.notify(GameEventSendMessage("Message received: ${payload.toString()}"))
        return state
    }
}