package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameEvent
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventObserver
import org.springframework.stereotype.Repository

@Repository
class InMemoryGameStreamCoordinator(
    private val observers: HashSet<GameStreamEventObserver> = HashSet()
):GameStreamEventCoordinator {
    override fun notify(event: GameEvent) {
        observers.forEach { observer -> observer.onEvent(event) }
    }

    override fun subscribe(observer: GameStreamEventObserver) {
        observers.add(observer)
    }

    override fun unsubscribe(observer: GameStreamEventObserver) {
        observers.remove(observer)
    }
}