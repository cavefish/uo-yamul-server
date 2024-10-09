package dev.cavefish.yamul.backend.game.controller.domain

data class MovementVector(val x: Int, val y: Int, val z: Int = 0) {
    fun add(v: MovementVector) = MovementVector(x + v.x, y + v.y, z + v.z)
    fun add(x: Int, y: Int, z: Int) = MovementVector(this.x + x, this.y + y, this.z + z)
}
