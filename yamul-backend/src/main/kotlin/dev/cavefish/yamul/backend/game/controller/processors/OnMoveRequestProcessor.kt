package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.common.api.Notoriety
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgClientMoveRequest
import dev.cavefish.yamul.backend.game.api.MsgMoveAck
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import org.springframework.stereotype.Component

@Component
class OnMoveRequestProcessor : MessageProcessor<MsgClientMoveRequest>(
    MsgType.TypeClientMoveRequest, Message::getClientMoveRequest
) {
    override fun process(
        payload: MsgClientMoveRequest,
        currentState: GameState,
        wrapper: GameStreamWrapper
    ): GameState {
        // TODO implement movement persistence and assertions
        wrapper.send(MsgType.TypeMoveAck) {
            it.setMoveAck(
                MsgMoveAck.newBuilder().setSequence(payload.sequence)
                    .setNotorietyFlagsValue(Notoriety.Unknown_VALUE)
            )
        }
        return currentState
    }
}