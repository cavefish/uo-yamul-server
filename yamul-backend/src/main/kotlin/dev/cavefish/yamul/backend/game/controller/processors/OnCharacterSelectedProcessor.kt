package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.common.api.Flags
import dev.cavefish.yamul.backend.common.api.Notoriety
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgApplyWorldPatches
import dev.cavefish.yamul.backend.game.api.MsgCharacterSelection
import dev.cavefish.yamul.backend.game.api.MsgExtendedStats
import dev.cavefish.yamul.backend.game.api.MsgGeneralLightLevel
import dev.cavefish.yamul.backend.game.api.MsgMapChange
import dev.cavefish.yamul.backend.game.api.MsgPlayMusic
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.MsgUpdateObject
import dev.cavefish.yamul.backend.game.api.MsgWarmode
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateErrorRequiresLoggedIn
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateLoggedIn
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimePosition
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import dev.cavefish.yamul.backend.game.controller.infra.UserCharacterRepository
import dev.cavefish.yamul.backend.game.controller.mappers.CharacterSkillUpdateMapper
import dev.cavefish.yamul.backend.game.controller.mappers.CharacterStatWindowMapper
import dev.cavefish.yamul.backend.game.controller.senders.PlayerStartConfirmationSender
import org.springframework.stereotype.Component
import org.tinylog.kotlin.Logger

@Component
class OnCharacterSelectedProcessor(
    private val playerStartConfirmationSender: PlayerStartConfirmationSender,
    private val userCharacterRepository: UserCharacterRepository,
    private val gameObjectRepository: GameObjectRepository,
    private val gameObjectRealtimePosition: GameObjectRealtimePosition,
    private val characterStatWindowMapper: CharacterStatWindowMapper,
    private val characterSkillUpdateMapper: CharacterSkillUpdateMapper,
) : MessageProcessor<MsgCharacterSelection>(MsgType.TypeCharacterSelection, Message::getCharacterSelection) {

    @SuppressWarnings("MaxLineLength", "MagicNumber", "LongMethod") // TODO remove exceptions
    override suspend fun process(payload: MsgCharacterSelection, state: State, wrapper: GameStreamWrapper): State {
        if (state !is StateLoggedIn) return StateErrorRequiresLoggedIn
        val character = userCharacterRepository.getCharacterByOrder(state.getLoggedUser(), payload.slot)!!
        val characterAsObject = gameObjectRepository.getById(character.objectId)!!
        val coordinatesOnRepo = gameObjectRealtimePosition.getCoordinates(character.objectId)
        val coordinates = if (coordinatesOnRepo != null) coordinatesOnRepo else {
            Logger.error("GameObject ${character.objectId} is not synchronized")
            Coordinates(
                x = 6787,
                y = 2181,
                z = 0
            )
        }
        val nextState = state.assignCharacter(
            characterObject = characterAsObject,
            coordinates = coordinates
        )
        playerStartConfirmationSender.send(nextState, wrapper)

        wrapper.send(MsgType.TypeApplyWorldPatches) {
            it.setApplyWorldPatches(MsgApplyWorldPatches.getDefaultInstance())
        }
        wrapper.send(MsgType.TypePlayMusic) { it.setPlayMusic(MsgPlayMusic.newBuilder().setMusicId(0x1E)) }
        wrapper.send(MsgType.TypeMapChange) {
            it.setMapChange(
                MsgMapChange.newBuilder().setMapId(coordinates.mapId)
            )
        }
        val gameObject = nextState.characterObject
        wrapper.send(MsgType.TypeUpdateObject) {
            it.setUpdateObject(
                MsgUpdateObject.newBuilder()
                    .setId(createObjectId(gameObject.id))
                    .setGraphicId(gameObject.graphicId.id)
                    .setHue(gameObject.hue.toUInt16().toInt())
                    .addAllFlags(gameObject.flags.map { f -> Flags.forNumber(f.id) })
                    .addAllNotorietyFlags(gameObject.notoriety.map { n -> Notoriety.forNumber(n.id) })
                    .addAllItems(gameObject.items.map { item ->
                        createMsgUpdateObjectItems(
                            item.id,
                            item.graphicId,
                            item.hue,
                            item.layer
                        ).build()
                    })
            )
        }
        wrapper.send(MsgType.TypeHealthBar) { it.setHealthBar(createMsgHealthBar(nextState.characterObject.id)) }
        wrapper.send(MsgType.TypeTeleportPlayer) { it.setTeleportPlayer(createMsgTeleportPlayer(nextState)) }
        wrapper.send(MsgType.TypeGeneralLightLevel) {
            it.setGeneralLightLevel(
                MsgGeneralLightLevel.newBuilder().setLevel(0x0)
            )
        }
        wrapper.send(MsgType.TypeStatWindowInfo) {
            it.setStatWindowInfo(
                characterStatWindowMapper.map(nextState.characterObject)
            )
        }
        wrapper.send(MsgType.TypeExtendedStats) {
            it.setExtendedStats(
                MsgExtendedStats.newBuilder().setLock(
                    MsgExtendedStats.MsgExtendedStats_AttributeLock.newBuilder()
                        .setId(createObjectId(nextState.characterObject.id))
                )
            )
        }
        wrapper.send(MsgType.TypeSkillUpdateServer) { it.setSkillUpdateServer(characterSkillUpdateMapper.getFullUpdate()) }
        wrapper.send(MsgType.TypeWarmode) { it.setWarmode(MsgWarmode.newBuilder().setIsWarmode(false)) }
        wrapper.send(MsgType.TypeLoginComplete) {}
        return nextState
    }


}
