package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgClientDoubleClick
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateErrorRequiresCharacter
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import dev.cavefish.yamul.backend.game.controller.mappers.OpenPaperDollMapper
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

private const val MASK = 0x7FFFFFFF

@Component
class OnClientDoubleClick(
    private val gameObjectRepository: GameObjectRepository,
    private val openPaperDollMapper: OpenPaperDollMapper,
) :
    MessageProcessor<MsgClientDoubleClick>(MsgType.TypeClientDoubleClick, Message::getClientDoubleClick) {
    override fun process(payload: MsgClientDoubleClick, state: State, wrapper: GameStreamWrapper): State {
        if (state !is StateHasCharacter) return StateErrorRequiresCharacter
        Logger.debug("Received onClientDoubleClick payload: $payload")

        val targetId =
            if (payload.target.value <= 0) state.characterObject.id else payload.target.value and MASK

        val gameObject = gameObjectRepository.getById(targetId)
        if (gameObject == null) {
            Logger.warn("Unknown objectId $targetId")
            return state
        }

        Logger.debug("Clicked on $gameObject")

        if (gameObject.id == state.characterObject.id) {
            Logger.debug("Click on self")
        }

        if (gameObject.isCharacter) {
            Logger.debug("Opening paper doll")
            wrapper.send(MsgType.TypeOpenPaperDoll) {it.setOpenPaperDoll(openPaperDollMapper.map(gameObject))}
        }

        return state
    }
}
