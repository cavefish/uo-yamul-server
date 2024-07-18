package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.common.api.Coordinate
import dev.cavefish.yamul.backend.common.api.Flags
import dev.cavefish.yamul.backend.common.api.Notoriety
import dev.cavefish.yamul.backend.common.api.ObjectId
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgApplyWorldPatches
import dev.cavefish.yamul.backend.game.api.MsgCharacterSelection
import dev.cavefish.yamul.backend.game.api.MsgExtendedStats
import dev.cavefish.yamul.backend.game.api.MsgGeneralLightLevel
import dev.cavefish.yamul.backend.game.api.MsgHealthBar
import dev.cavefish.yamul.backend.game.api.MsgMapChange
import dev.cavefish.yamul.backend.game.api.MsgPlayMusic
import dev.cavefish.yamul.backend.game.api.MsgTeleportPlayer
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.MsgUpdateObject
import dev.cavefish.yamul.backend.game.api.MsgUpdateObjectItems
import dev.cavefish.yamul.backend.game.api.MsgWarmode
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.domain.GraphicId
import dev.cavefish.yamul.backend.game.controller.domain.Hue
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRealtimeLocalization
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
    private val gameObjectRealtimeLocalization: GameObjectRealtimeLocalization,
    private val characterStatWindowMapper: CharacterStatWindowMapper,
    private val characterSkillUpdateMapper: CharacterSkillUpdateMapper,
) : MessageProcessor<MsgCharacterSelection>(MsgType.TypeCharacterSelection, Message::getCharacterSelection) {

    @SuppressWarnings("MaxLineLength", "MagicNumber", "LongMethod") // TODO remove exceptions
    override fun process(
        payload: MsgCharacterSelection,
        currentState: GameState?,
        loggedUser: LoggedUser,
        wrapper: GameStreamWrapper
    ): GameState {
        val character = userCharacterRepository.getCharacterByOrder(loggedUser, payload.slot)!!
        val characterAsObject = gameObjectRepository.getById(character.objectId)!!
        val coordinatesOnRepo = gameObjectRealtimeLocalization.getCoordinates(character.objectId)
        val coordinates = if (coordinatesOnRepo != null) coordinatesOnRepo else {
            Logger.error("GameObject ${character.objectId} is not synchronized")
            Coordinates(
                x = 6787,
                y = 2181,
                z = 0
            )
        }
        val nextState = GameState(
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
                MsgMapChange.newBuilder().setMapId(1)
            )
        }
        val gameObject = nextState.characterObject
        wrapper.send(MsgType.TypeUpdateObject) {
            it.setUpdateObject(
                MsgUpdateObject.newBuilder()
                    .setId(createObjectId(gameObject.id))
                    .setGraphicId(gameObject.graphicId.id)
                    .setHue(gameObject.hue.toInt16())
                    .addAllFlags(gameObject.flags.map { f -> Flags.forNumber(f.id) })
                    .addAllNotorietyFlags(gameObject.notoriety.map { n-> Notoriety.forNumber(n.id) })
                    .addAllItems(gameObject.items.map { item ->
                        createItem(
                            item.id,
                            item.graphicId,
                            item.hue,
                            item.layer
                        ).build()
                    })
            )
        }
        wrapper.send(MsgType.TypeHealthBar) { it.setHealthBar(createHealthBar(nextState.characterObject.id)) }
        wrapper.send(MsgType.TypeTeleportPlayer) { it.setTeleportPlayer(createTeleportPlayer(nextState)) }
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


    private fun createItem(id: Int, graphicId: GraphicId, hue: Hue, layer: Int): MsgUpdateObjectItems.Builder =
        MsgUpdateObjectItems.newBuilder().setId(createObjectId(id)).setGraphicId(graphicId.id)
            .setHue(hue.toInt16()).setLayer(layer)


    private fun createHealthBar(objectId: Int): MsgHealthBar.Builder = MsgHealthBar.newBuilder()
        .setId(createObjectId(objectId))
        .addValues(
            MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.GREEN_VALUE).setEnabled(false)
        )
        .addValues(
            MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.YELLOW_VALUE).setEnabled(false)
        )


    private fun createTeleportPlayer(state: GameState): MsgTeleportPlayer.Builder =
        MsgTeleportPlayer.newBuilder()
            .setId(createObjectId(state.characterObject.id))
            .setCoordinates(createPlayerObjectCoordinates(state))
            .setGraphicId(state.characterObject.graphicId.id)
            .setHue(state.characterObject.hue.toInt16())
            .addAllStatusValue(state.characterObject.flags.map { f -> f.id })

    private fun createPlayerObjectCoordinates(state: GameState): Coordinate.Builder =
        Coordinate.newBuilder().setXLoc(state.coordinates.x).setYLoc(state.coordinates.y).setZLoc(state.coordinates.z)

    private fun createObjectId(objectId: Int): ObjectId.Builder = ObjectId.newBuilder().setValue(objectId)
}
