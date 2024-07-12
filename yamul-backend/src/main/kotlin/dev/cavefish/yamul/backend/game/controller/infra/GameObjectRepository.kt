package dev.cavefish.yamul.backend.game.controller.infra

import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId

interface GameObjectRepository {
    fun getById(id: ObjectId): GameObject?
}