package dev.cavefish.yamul.backend.game.controller.mappers

import dev.cavefish.yamul.backend.common.api.ObjectId
import org.springframework.stereotype.Service

@Service
class ObjectIdMapper {
    fun create(value: Int) = ObjectId.newBuilder().setValue(value)
}