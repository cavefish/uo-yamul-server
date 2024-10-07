package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.MsgUnicodeSpeechSelected
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.UserMessage
import dev.cavefish.yamul.backend.game.controller.domain.UserMessageType
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter
import dev.cavefish.yamul.backend.game.controller.infra.MessageCoordinator
import org.springframework.stereotype.Component

@Component
class OnUnicodeSpeechSelected(
    private val messageCoordinator: MessageCoordinator
) : MessageProcessor<MsgUnicodeSpeechSelected>(
    MsgType.TypeUnicodeSpeechSelected, Message::getUnicodeSpeechSelected
) {

    override fun process(payload: MsgUnicodeSpeechSelected, state: State, wrapper: GameStreamWrapper): State {
        if (state is StateHasCharacter) messageCoordinator.onUserMessage(
            UserMessage(
                originObject = state.characterObject,
                text = payload.text,
                type = UserMessageType.mapFromApi(payload.mode)
            )
            , state)
        return state
    }
}