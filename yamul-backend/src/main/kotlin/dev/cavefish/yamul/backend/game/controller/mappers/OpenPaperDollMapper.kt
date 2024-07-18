package dev.cavefish.yamul.backend.game.controller.mappers

import dev.cavefish.yamul.backend.game.api.MsgOpenPaperDoll
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import org.springframework.stereotype.Service

@Service
class OpenPaperDollMapper(
    private val objectIdMapper: ObjectIdMapper
) {
    fun map(character: GameObject) : MsgOpenPaperDoll.Builder = MsgOpenPaperDoll.newBuilder()
        .setId(objectIdMapper.create(character.id))
        .setName(character.name)
        .setFlags(character.flags)
}