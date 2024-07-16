package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgClientDoubleClick
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

@Component
class OnClientDoubleClick :
    MessageProcessor<MsgClientDoubleClick>(MsgType.TypeClientDoubleClick, Message::getClientDoubleClick) {
    override fun process(
        payload: MsgClientDoubleClick,
        currentState: GameState?,
        loggedUser: LoggedUser,
        wrapper: GameStreamWrapper
    ): GameState? {
        Logger.info("Received onClientDoubleClick payload: $payload")
        return currentState
    }
}
