package dev.cavefish.yamul.backend.game.controller.domain.gameevents

import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameState

sealed class GameEvent(open val filter: GameEventFilter) {
    fun appliesTo(state: GameState) = filter.invoke(state)
    abstract operator fun invoke(state: GameState, streamWrapper: GameStreamWrapper)
}

