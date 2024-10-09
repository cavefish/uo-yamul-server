package dev.cavefish.yamul.backend.game.controller.domain

data class Coordinates(
    val x: Int,
    val y: Int,
    val z: Int = 0,
    val mapId: Int = -1
) {
    fun collidesWith(coordinates: Coordinates): Boolean =
        x == coordinates.x && y == coordinates.y && mapId == coordinates.mapId

    fun applyMovement(movement: MovementVector): Coordinates = this.copy(
        x = movement.x + x,
        y = movement.y + y,
        z = movement.z + z
    )

    fun difference(other: Coordinates) = Coordinates(
        this.x - other.x,
        this.y - other.y,
        this.z - other.z
    )

    fun toBlockOrigin(): Coordinates = Coordinates(
        x = roundToLowerDivisorOf8(this.x),
        y = roundToLowerDivisorOf8(this.y),
        z = 0,
        mapId = this.mapId
    )

    @SuppressWarnings("MagicNumber")
    private fun roundToLowerDivisorOf8(x: Int) = x and 0xFFF8
}
