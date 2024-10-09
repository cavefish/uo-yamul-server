package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimePosition
import org.springframework.stereotype.Repository
import java.util.concurrent.ConcurrentHashMap
import java.util.concurrent.locks.ReentrantLock
import kotlin.concurrent.withLock

@Repository
class InMemoryGameObjectRealtimePosition : GameObjectRealtimePosition {
    override suspend fun getCoordinates(id: ObjectId): Coordinates? = mutex.withLock {
        return allCoordinates[id]
    }

    override suspend fun registerNewCoordinates(id: ObjectId, coordinates: Coordinates): Boolean = mutex.withLock {
        if (allCoordinates.containsKey(id)) return false
        if (allCoordinates.values.any(coordinates::collidesWith)) return false
        allCoordinates[id] = coordinates
        return true
    }

    override suspend fun areCoordinatesEmpty(coordinates: Coordinates): Boolean = mutex.withLock {
        return allCoordinates.values.none(coordinates::collidesWith)
    }

    override suspend fun updatePosition(id: ObjectId, movement: (Coordinates) -> Coordinates): Coordinates? =
        mutex.withLock {
            val oldCoordinate = allCoordinates[id]
            val nextCoordinate = movement(oldCoordinate!!)
            if (allCoordinates.values.any {
                    it != oldCoordinate && nextCoordinate.collidesWith(it)
                }) return null
            allCoordinates[id] = nextCoordinate
            return@withLock nextCoordinate
        }

    private companion object {
        val allCoordinates = ConcurrentHashMap<ObjectId, Coordinates>()
        val mutex = ReentrantLock()
    }
}