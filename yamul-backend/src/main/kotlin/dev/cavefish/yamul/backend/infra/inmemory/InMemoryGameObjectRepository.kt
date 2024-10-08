package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.GameObjectItem
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import org.springframework.stereotype.Repository
import java.util.concurrent.ConcurrentHashMap
import java.util.concurrent.ConcurrentLinkedQueue
import java.util.concurrent.atomic.AtomicInteger

@Repository
class InMemoryGameObjectRepository(
    private val counter: AtomicInteger = AtomicInteger(1)
):GameObjectRepository {
    override fun getById(id: ObjectId): GameObject? {
        val gameObject = database[id]
        val children = childrenIndex[id]
        if (children.isNullOrEmpty()) {
            return gameObject
        }
        return gameObject?.copy(items = children.mapNotNull(this::toGameObjectItem))
    }

    private fun toGameObjectItem(id: ObjectId): GameObjectItem? {
        val gameObject = database[id] ?: return null
        return GameObjectItem(
            id = id,
            graphicId = gameObject.graphicId,
            hue = gameObject.hue,
            layer = gameObject.layer,
        )
    }

    override fun registerNewObject(newObject: GameObject): ObjectId {
        val id = counter.getAndIncrement()
        database[id] = newObject.copy(id = id, items = emptyList())
        if (newObject.parentId != null) {
            childrenIndex.getOrPut(newObject.parentId) {ConcurrentLinkedQueue<ObjectId>()}.add(id)
        }
        return id
    }

    override fun updateObject(id: ObjectId, updateAction: (GameObject) -> GameObject): GameObject {
        return database.compute(id) { _, obj ->
            updateAction(obj!!)
        }!!
    }

    private companion object {
        val childrenIndex = ConcurrentHashMap<ObjectId, ConcurrentLinkedQueue<ObjectId>>()
        val database = ConcurrentHashMap<ObjectId, GameObject>()
    }
}
