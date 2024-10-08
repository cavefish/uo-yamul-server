package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.common.api.Coordinate
import dev.cavefish.yamul.backend.common.api.ObjectId
import dev.cavefish.yamul.backend.game.api.MsgHealthBar
import dev.cavefish.yamul.backend.game.api.MsgTeleportPlayer
import dev.cavefish.yamul.backend.game.api.MsgUpdateObjectItems
import dev.cavefish.yamul.backend.game.controller.domain.GraphicId
import dev.cavefish.yamul.backend.game.controller.domain.Hue
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.StateHasCharacter

fun createObjectId(objectId: Int): ObjectId.Builder = ObjectId.newBuilder().setValue(objectId)

fun createMsgTeleportPlayer(state: StateHasCharacter): MsgTeleportPlayer.Builder =
    MsgTeleportPlayer.newBuilder()
        .setId(createObjectId(state.characterObject.id))
        .setCoordinates(createPlayerObjectCoordinates(state))
        .setDirectionValue(state.characterObject.facing.apiValue.number)
        .setGraphicId(state.characterObject.graphicId.id)
        .setHue(state.characterObject.hue.toUInt16().toInt())
        .addAllStatusValue(state.characterObject.flags.map { f -> f.id })


fun createPlayerObjectCoordinates(state: StateHasCharacter): Coordinate.Builder =
    Coordinate.newBuilder().setXLoc(state.coordinates.x).setYLoc(state.coordinates.y).setZLoc(state.coordinates.z)

fun createMsgUpdateObjectItems(id: Int, graphicId: GraphicId, hue: Hue, layer: Int): MsgUpdateObjectItems.Builder =
    MsgUpdateObjectItems.newBuilder().setId(createObjectId(id)).setGraphicId(graphicId.id)
        .setHue(hue.toUInt16().toInt()).setLayer(layer)


fun createMsgHealthBar(objectId: Int): MsgHealthBar.Builder = MsgHealthBar.newBuilder()
    .setId(createObjectId(objectId))
    .addValues(
        MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.GREEN_VALUE).setEnabled(false)
    )
    .addValues(
        MsgHealthBar.Values.newBuilder().setTypeValue(MsgHealthBar.Values.Type.YELLOW_VALUE).setEnabled(false)
    )
