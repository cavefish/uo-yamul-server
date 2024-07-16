package dev.cavefish.yamul.backend.game.controller.senders

import dev.cavefish.yamul.backend.game.api.MsgPlayerStartConfirmation
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.mappers.CoordinateMapper
import dev.cavefish.yamul.backend.game.controller.mappers.ObjectIdMapper
import org.springframework.stereotype.Service

@Service
class PlayerStartConfirmationSender (
    private val objectIdMapper: ObjectIdMapper,
    private val coordinateMapper: CoordinateMapper
) {
    fun send(state: GameState, wrapper: GameStreamWrapper) {
        wrapper.send(MsgType.TypePlayerStartConfirmation) {
            it.setPlayerStartConfirmation(
                MsgPlayerStartConfirmation.newBuilder()
                    .setId(objectIdMapper.create(state.characterObject.id))
                    .setCoordinates(coordinateMapper.map(state.coordinates))
                    .setGraphicId(state.characterObject.graphicId.id)
                    .setHue(state.characterObject.hue.toInt16())
            )
        }
    }
}
