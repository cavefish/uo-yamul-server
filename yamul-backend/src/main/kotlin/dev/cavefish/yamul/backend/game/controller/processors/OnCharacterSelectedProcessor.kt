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
import dev.cavefish.yamul.backend.game.api.MsgWarmode
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.senders.PlayerStartConfirmationSender
import org.springframework.stereotype.Component

@Component
class OnCharacterSelectedProcessor (
    private val playerStartConfirmationSender: PlayerStartConfirmationSender
) : MessageProcessor<MsgCharacterSelection>(MsgType.TypeCharacterSelection, Message::getCharacterSelection) {

    override fun process(
        payload: MsgCharacterSelection,
        currentState: GameState,
        wrapper: GameStreamWrapper
    ): GameState {
        val nextState = currentState.copy(
            characterObjectId = 1,
            coordinates = Coordinates(x = 6787, y = 2181, z = 0)
        )
        playerStartConfirmationSender.send(nextState, wrapper)

        wrapper.send(MsgType.TypeApplyWorldPatches) {
            it.setApplyWorldPatches(MsgApplyWorldPatches.getDefaultInstance())
        }
        wrapper.send(MsgType.TypePlayMusic) { it.setPlayMusic(MsgPlayMusic.newBuilder().setMusicId(0x1E)) }
        wrapper.send(MsgType.TypeMapChange) {
            it.setMapChange(
                MsgMapChange.newBuilder().setMapId(nextState.characterObjectId)
            )
        }
        wrapper.send(MsgType.TypeUpdateObject) {
            it.setUpdateObject(
                MsgUpdateObject.newBuilder()
                    .setId(createPlayerObjectId(nextState.characterObjectId))
                    .setGraphicId(0x0190)
                    .setHue(0x83EA)
            )
        }
        wrapper.send(MsgType.TypeHealthBar) { it.setHealthBar(createHealthBar(nextState.characterObjectId)) }
        wrapper.send(MsgType.TypeTeleportPlayer) { it.setTeleportPlayer(createTeleportPlayer()) }
        wrapper.send(MsgType.TypeGeneralLightLevel) {
            it.setGeneralLightLevel(
                MsgGeneralLightLevel.newBuilder().setLevel(0x18)
            )
        }
        wrapper.send(MsgType.TypeStatWindowInfo) {
            it.setStatWindowInfo(
                MsgStatWindowInfo.newBuilder().setCharacterID(createPlayerObjectId(nextState.characterObjectId))
            )
        }
        wrapper.send(MsgType.TypeExtendedStats) { it.setExtendedStats(MsgExtendedStats.getDefaultInstance()) }
        wrapper.send(MsgType.TypeWarmode) { it.setWarmode(MsgWarmode.getDefaultInstance()) }
        wrapper.send(MsgType.TypeLoginComplete) {}
        return nextState
    }


    private fun createHealthBar(objectId: Int): MsgHealthBar.Builder = MsgHealthBar.newBuilder()
        .setId(createPlayerObjectId(objectId))
        .addValues(MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.GREEN_VALUE).setEnabled(true))
        .addValues(
            MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.YELLOW_VALUE).setEnabled(false)
        )


    private fun createTeleportPlayer(): MsgTeleportPlayer.Builder =
        MsgTeleportPlayer.newBuilder()
            .setId(createPlayerObjectId(1))
            .setCoordinates(createPlayerObjectCoordinates())

    private fun createPlayerObjectCoordinates(): Coordinate.Builder =
        Coordinate.newBuilder().setXLoc(6787).setYLoc(2181).setZLoc(0)

    private fun createPlayerObjectId(objectId: Int): ObjectId.Builder = ObjectId.newBuilder().setValue(objectId)
}