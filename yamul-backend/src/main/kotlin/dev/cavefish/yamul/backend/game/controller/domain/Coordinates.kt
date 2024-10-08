package dev.cavefish.yamul.backend.game.controller.domain

data class Coordinates(
    val x: Int = 0,
    val y: Int = 0,
    val z: Int = 0,
    val mapId: Int = 0
) {
    fun collidesWith(coordinates: Coordinates): Boolean =
        x == coordinates.x && y == coordinates.y && mapId == coordinates.mapId
}
