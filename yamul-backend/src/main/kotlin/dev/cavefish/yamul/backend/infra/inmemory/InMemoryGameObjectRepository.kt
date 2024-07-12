package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import org.springframework.stereotype.Repository

@Repository
class InMemoryGameObjectRepository:GameObjectRepository {
    override fun getById(id: ObjectId): GameObject? = database[id]

    companion object {
        val database = HashMap<ObjectId, GameObject>()
    }
}