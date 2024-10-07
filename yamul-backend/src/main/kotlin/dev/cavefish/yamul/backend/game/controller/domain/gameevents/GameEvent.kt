package dev.cavefish.yamul.backend.game.controller.domain.gameevents

import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State

sealed class GameEvent(open val filter: GameEventFilter) {
    fun appliesTo(state: State) = filter.invoke(state)
    abstract operator fun invoke(state: State, streamWrapper: GameStreamWrapper)
}

