package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.common.api.Coordinate
import dev.cavefish.yamul.backend.common.api.ObjectId
import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgApplyWorldPatches
import dev.cavefish.yamul.backend.game.api.MsgCharacterSelection
import dev.cavefish.yamul.backend.game.api.MsgExtendedStats
import dev.cavefish.yamul.backend.game.api.MsgGeneralLightLevel
import dev.cavefish.yamul.backend.game.api.MsgHealthBar
import dev.cavefish.yamul.backend.game.api.MsgMapChange
import dev.cavefish.yamul.backend.game.api.MsgPlayMusic
import dev.cavefish.yamul.backend.game.api.MsgStatWindowInfo
import dev.cavefish.yamul.backend.game.api.MsgTeleportPlayer
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.MsgUpdateObject
import dev.cavefish.yamul.backend.game.api.MsgUpdateObjectItems
import dev.cavefish.yamul.backend.game.api.MsgWarmode
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.Flags
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.domain.GraphicIds
import dev.cavefish.yamul.backend.game.controller.domain.Hue
import dev.cavefish.yamul.backend.game.controller.domain.Hues
import dev.cavefish.yamul.backend.game.controller.domain.Notoriety
import dev.cavefish.yamul.backend.game.controller.senders.PlayerStartConfirmationSender
import org.springframework.stereotype.Component

@Component
class OnCharacterSelectedProcessor (
    private val playerStartConfirmationSender: PlayerStartConfirmationSender
) : MessageProcessor<MsgCharacterSelection>(MsgType.TypeCharacterSelection, Message::getCharacterSelection) {

    @SuppressWarnings("MaxLineLength")
    override fun process(
        payload: MsgCharacterSelection,
        currentState: GameState,
        wrapper: GameStreamWrapper
    ): GameState {
        val nextState = currentState.copy(
            characterObjectId = 1,
            characterBodyType = GraphicIds.BodyHumanMale,
            characterBodyHue = Hues.White.hue,
            coordinates = Coordinates(x = 6787, y = 2181, z = 0)
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
        wrapper.send(MsgType.TypeUpdateObject) {
            it.setUpdateObject(
                MsgUpdateObject.newBuilder()
                    .setId(createObjectId(nextState.characterObjectId))
                    .setGraphicId(nextState.characterBodyType.id)
                    .setHue(nextState.characterBodyHue.toInt16())
                    .setFlags(Flags.None.id)
                    .setNotorietyFlagsValue(Notoriety.Innocent.id)
                    .addItems(createItem(2, GraphicIds.Backpack, Hues.Blue.hue, 0x15))
                    .addItems(createItem(3, GraphicIds.HairShort, Hues.Green.hue, 0x0B))
                    .addItems(createItem(4, GraphicIds.RobeGm, Hues.White.hue, 0x16))
            )
        }
        wrapper.send(MsgType.TypeHealthBar) { it.setHealthBar(createHealthBar(nextState.characterObjectId)) }
        wrapper.send(MsgType.TypeTeleportPlayer) { it.setTeleportPlayer(createTeleportPlayer(nextState)) }
        wrapper.send(MsgType.TypeGeneralLightLevel) {
            it.setGeneralLightLevel(
                MsgGeneralLightLevel.newBuilder().setLevel(0x18)
            )
        }
        wrapper.send(MsgType.TypeStatWindowInfo) {
            it.setStatWindowInfo(
                MsgStatWindowInfo.newBuilder().setCharacterID(createObjectId(nextState.characterObjectId))
            )
        }
        wrapper.send(MsgType.TypeExtendedStats) { it.setExtendedStats(MsgExtendedStats.getDefaultInstance()) }
        wrapper.send(MsgType.TypeWarmode) { it.setWarmode(MsgWarmode.getDefaultInstance()) }
        wrapper.send(MsgType.TypeLoginComplete) {}
        return nextState
    }

    private fun createItem(id: Int, graphicId: GraphicIds, hue: Hue, layer: Int): MsgUpdateObjectItems.Builder? =
        MsgUpdateObjectItems.newBuilder().setId(createObjectId(id)).setGraphicId(graphicId.id)
            .setHue(hue.toInt16()).setLayer(layer)


    private fun createHealthBar(objectId: Int): MsgHealthBar.Builder = MsgHealthBar.newBuilder()
        .setId(createObjectId(objectId))
        .addValues(MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.GREEN_VALUE).setEnabled(true))
        .addValues(
            MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.YELLOW_VALUE).setEnabled(false)
        )


    private fun createTeleportPlayer(state: GameState): MsgTeleportPlayer.Builder =
        MsgTeleportPlayer.newBuilder()
            .setId(createObjectId(state.characterObjectId))
            .setCoordinates(createPlayerObjectCoordinates(state))
            .setGraphicId(state.characterBodyType.id)
            .setHue(state.characterBodyHue.toInt16())
            .addStatusValue(0) // TODO remove hardcoded value

    private fun createPlayerObjectCoordinates(state: GameState): Coordinate.Builder =
        Coordinate.newBuilder().setXLoc(state.coordinates.x).setYLoc(state.coordinates.y).setZLoc(state.coordinates.z)

    private fun createObjectId(objectId: Int): ObjectId.Builder = ObjectId.newBuilder().setValue(objectId)
}