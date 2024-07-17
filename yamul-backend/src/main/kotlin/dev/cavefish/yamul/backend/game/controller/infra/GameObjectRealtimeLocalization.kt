package dev.cavefish.yamul.backend.game.controller.infra

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId

interface GameObjectRealtimeLocalization {
    fun getCoordinates(id: ObjectId): Coordinates?
}