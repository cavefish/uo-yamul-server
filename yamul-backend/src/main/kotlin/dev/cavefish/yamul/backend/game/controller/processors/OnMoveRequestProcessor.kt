package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.Constants.HUMAN_BODY_HEIGTH
import dev.cavefish.yamul.backend.common.api.Notoriety
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgClientMoveRequest
import dev.cavefish.yamul.backend.game.api.MsgMoveAck
import dev.cavefish.yamul.backend.game.api.MsgMoveReject
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.MovementDirection
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateError
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateErrorRequiresCharacter
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimePosition
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulMapBlockRepository
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

private const val RESYNC_SEQUENCE = 5

@Component
class OnMoveRequestProcessor(
    private val gameObjectRepository: GameObjectRepository,
    private val gameObjectRealtimePosition: GameObjectRealtimePosition,
    private val mulMapBlockRepository: MulMapBlockRepository
) : MessageProcessor<MsgClientMoveRequest>(
    MsgType.TypeClientMoveRequest, Message::getClientMoveRequest
) {
    override suspend fun process(payload: MsgClientMoveRequest, state: State, wrapper: GameStreamWrapper): State {
        if (state !is StateHasCharacter) return StateErrorRequiresCharacter
        val movementFacing = MovementDirection.fromApi(payload.direction!!)
            ?: return StateError("Unknown movement direction ${payload.direction}")

        Logger.debug(payload)

        val sameFacing = state.characterObject.facing == movementFacing
        val newCoordinates = if (sameFacing) {
            tryToUpdatePosition(state, movementFacing) ?: return processCharacterCollision(state, wrapper, payload)
        } else {
            state.coordinates
        }

        val updatedCharacterObject = gameObjectRepository.updateObject(state.characterObject.id) {
            it.copy(facing = movementFacing)
        }

        val nextState = state.copy(characterObject = updatedCharacterObject, coordinates = newCoordinates)

        if (payload.sequence == RESYNC_SEQUENCE) {
            sendMoveReject(wrapper, nextState, payload.sequence)
        } else {
            sendMoveAck(wrapper, payload)
        }

        return nextState
    }

    private suspend fun processCharacterCollision(
        state: StateHasCharacter,
        wrapper: GameStreamWrapper,
        payload: MsgClientMoveRequest
    ): StateHasCharacter {
        Logger.warn("Character collision")
        val objectId = state.characterObject.id
        val coordinates = gameObjectRealtimePosition.getCoordinates(objectId) ?: state.coordinates
        val gameObject = gameObjectRepository.getById(objectId) ?: state.characterObject
        val nextState = state.copy(coordinates = coordinates, characterObject = gameObject)
        sendMoveReject(wrapper, nextState, payload.sequence)
        return nextState
    }

    private suspend fun tryToUpdatePosition(
        state: StateHasCharacter,
        movementFacing: MovementDirection
    ) = gameObjectRealtimePosition.updatePosition(state.characterObject.id) {
        mulMapBlockRepository.correctPositionAltitude(
            it.applyMovement(
                movementFacing.movement
            ),
            HUMAN_BODY_HEIGTH
        )
    }

    private fun sendMoveAck(
        wrapper: GameStreamWrapper,
        payload: MsgClientMoveRequest
    ) {
        wrapper.send(MsgType.TypeMoveAck) {
            it.setMoveAck(
                MsgMoveAck.newBuilder().setSequence(payload.sequence)
                    .setNotorietyFlagsValue(Notoriety.Unknown_VALUE)
            )
        }
    }

    private fun sendMoveReject(
        wrapper: GameStreamWrapper,
        state: StateHasCharacter,
        sequence: Int
    ) {
        wrapper.send(MsgType.TypeMoveReject) {
            it.setMoveReject(
                MsgMoveReject.newBuilder()
                    .setSequence(sequence)
                    .setXLoc(state.coordinates.x)
                    .setYLoc(state.coordinates.y)
                    .setZLoc(state.coordinates.z)
                    .setDirectionValue(state.characterObject.facing.apiValue.number)
            )
        }

        wrapper.send(MsgType.TypeTeleportPlayer) { it.setTeleportPlayer(createMsgTeleportPlayer(state)) }
    }

}
