package dev.cavefish.yamul.backend.game.controller.domain.gameevents

interface GameStreamEventCoordinator {
    fun notify(event: GameEvent)
    fun subscribe(observer: GameStreamEventObserver)
    fun unsubscribe(observer: GameStreamEventObserver)
}