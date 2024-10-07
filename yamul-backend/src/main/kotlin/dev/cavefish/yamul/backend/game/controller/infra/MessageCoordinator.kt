package dev.cavefish.yamul.backend.game.controller.infra

import dev.cavefish.yamul.backend.game.controller.domain.UserMessage
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter

interface MessageCoordinator {
    fun onUserMessage(msg: UserMessage, state: StateHasCharacter)
}