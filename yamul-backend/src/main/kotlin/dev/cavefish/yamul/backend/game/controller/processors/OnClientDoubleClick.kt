package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgClientDoubleClick
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

@Component
class OnClientDoubleClick(
    private val gameObjectRepository: GameObjectRepository
) :
    MessageProcessor<MsgClientDoubleClick>(MsgType.TypeClientDoubleClick, Message::getClientDoubleClick) {
    override fun process(
        payload: MsgClientDoubleClick,
        currentState: GameState?,
        loggedUser: LoggedUser,
        wrapper: GameStreamWrapper
    ): GameState? {
        Logger.debug("Received onClientDoubleClick payload: $payload")

        val targetId = payload.target.value and 0x7FFFFFFF

        val gameObject = gameObjectRepository.getById(targetId)
        if (gameObject == null) {
            Logger.warn("Unknown objectId $targetId")
            return currentState
        }

        Logger.info("Clicked on $gameObject")

        return currentState
    }
}
