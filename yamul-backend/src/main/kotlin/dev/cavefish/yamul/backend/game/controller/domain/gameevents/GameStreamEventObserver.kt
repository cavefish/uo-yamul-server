package dev.cavefish.yamul.backend.game.controller.domain.gameevents

interface GameStreamEventObserver {
    fun onEvent(event: GameEvent)
}