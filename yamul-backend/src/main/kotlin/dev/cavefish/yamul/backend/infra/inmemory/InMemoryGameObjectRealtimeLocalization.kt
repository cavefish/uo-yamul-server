package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimeLocalization
import org.springframework.stereotype.Repository

@Repository
class InMemoryGameObjectRealtimeLocalization: GameObjectRealtimeLocalization {
    override fun getCoordinates(id: ObjectId): Coordinates = Coordinates(
        x = 6787,
        y = 2181,
        z = 0,
        mapId = 1,
    )
}