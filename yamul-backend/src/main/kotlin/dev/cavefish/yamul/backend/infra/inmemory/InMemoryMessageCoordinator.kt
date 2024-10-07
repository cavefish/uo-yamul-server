package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.UserMessage
import dev.cavefish.yamul.backend.game.controller.domain.UserMessageType
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameEventSendMessage
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter
import dev.cavefish.yamul.backend.game.controller.infra.MessageCoordinator
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

@Component
class InMemoryMessageCoordinator(
    private val eventCoordinator: GameStreamEventCoordinator
): MessageCoordinator {
    override fun onUserMessage(msg: UserMessage, state: StateHasCharacter) {
        if (msg.type!=UserMessageType.Normal) {
            Logger.warn("Unexpected message type ${msg.type}")
            return
        }
        eventCoordinator.notify(GameEventSendMessage(
            message = msg.text,
            type = msg.type,
            origin = msg.originObject
        ))
    }
}