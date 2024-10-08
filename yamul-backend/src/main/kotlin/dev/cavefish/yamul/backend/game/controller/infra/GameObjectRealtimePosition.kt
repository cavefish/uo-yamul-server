package dev.cavefish.yamul.backend.game.controller.infra

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId

interface GameObjectRealtimePosition {
    suspend fun getCoordinates(id: ObjectId): Coordinates?
    suspend fun registerNewCoordinates(id: ObjectId, coordinates: Coordinates): Boolean
    suspend fun areCoordinatesEmpty(coordinates: Coordinates): Boolean
}