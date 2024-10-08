package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.common.api.Notoriety
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgClientMoveRequest
import dev.cavefish.yamul.backend.game.api.MsgMoveAck
import dev.cavefish.yamul.backend.game.api.MsgTeleportPlayer
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.MovementDirection
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateError
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateErrorRequiresCharacter
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimePosition
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

private const val RESYNC_SEQUENCE = 250

@Component
class OnMoveRequestProcessor(
    private val gameObjectRepository: GameObjectRepository,
    private val gameObjectRealtimePosition: GameObjectRealtimePosition
) : MessageProcessor<MsgClientMoveRequest>(
    MsgType.TypeClientMoveRequest, Message::getClientMoveRequest
) {
    override suspend fun process(payload: MsgClientMoveRequest, state: State, wrapper: GameStreamWrapper): State {
        if (state !is StateHasCharacter) return StateErrorRequiresCharacter
        val movementFacing = MovementDirection.fromApi(payload.direction!!)
            ?: return StateError("Unknown movement direction ${payload.direction}")

        val newCoordinates = if (state.characterObject.facing == movementFacing) {
            val updatedValue =
                gameObjectRealtimePosition.updatePosition(state.characterObject.id, movementFacing.movement)
            if (updatedValue == null) {
                Logger.warn("Character collision")
                wrapper.send(MsgType.TypeTeleportPlayer) {
                    it.setTeleportPlayer(
                        createMsgTeleportPlayer(state)
                    )
                }
                return state
            }
            updatedValue
        } else {
            state.coordinates
        }

        val updatedCharacterObject = gameObjectRepository.updateObject(state.characterObject.id) {
            it.copy(facing = movementFacing)
        }

        val nextState = state.copy(characterObject = updatedCharacterObject, coordinates = newCoordinates)



        if (payload.sequence == RESYNC_SEQUENCE) {
            wrapper.send(MsgType.TypeTeleportPlayer) {
                it.setTeleportPlayer(
                    createMsgTeleportPlayer(nextState)
                )
            }
        } else {
            wrapper.send(MsgType.TypeMoveAck) {
                it.setMoveAck(
                    MsgMoveAck.newBuilder().setSequence(payload.sequence)
                        .setNotorietyFlagsValue(Notoriety.Unknown_VALUE)
                )
            }
        }

        return nextState
    }

}
