package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimePosition
import org.springframework.stereotype.Repository
import java.util.*
import java.util.concurrent.locks.ReentrantLock
import kotlin.concurrent.withLock

@Repository
class InMemoryGameObjectRealtimePosition : GameObjectRealtimePosition {
    override suspend fun getCoordinates(id: ObjectId): Coordinates? = mutex.withLock {
        return data[id]
    }

    override suspend fun registerNewCoordinates(id: ObjectId, coordinates: Coordinates): Boolean = mutex.withLock {
        if (data.containsKey(id)) return false
        if (data.values.any(coordinates::collidesWith)) return false
        data[id] = coordinates
        return true
    }

    override suspend fun areCoordinatesEmpty(coordinates: Coordinates): Boolean = mutex.withLock {
        return data.values.none(coordinates::collidesWith)
    }

    private companion object {
        val data: MutableMap<ObjectId, Coordinates> = Collections.synchronizedMap(HashMap())
        val mutex = ReentrantLock()
    }
}