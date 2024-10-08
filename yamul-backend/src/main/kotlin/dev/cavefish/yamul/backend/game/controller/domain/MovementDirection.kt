package dev.cavefish.yamul.backend.game.controller.domain

import dev.cavefish.yamul.backend.common.api.ObjectDirection

enum class MovementDirection(val movement: MovementVector, val apiValue: ObjectDirection) {

    North(MovementVector(0, -1), ObjectDirection.north),
    Right(MovementVector(1, -1), ObjectDirection.right),
    East(MovementVector(1, 0), ObjectDirection.east),
    Down(MovementVector(1, 1), ObjectDirection.down),
    South(MovementVector(0, 1), ObjectDirection.south),
    Left(MovementVector(-1, 1), ObjectDirection.left),
    West(MovementVector(-1, 0), ObjectDirection.west),
    Up(MovementVector(-1, -1), ObjectDirection.up), ;

    companion object {
        @SuppressWarnings("CyclomaticComplexMethod")
        fun fromApi(direction: ObjectDirection): MovementDirection? {
            return when (direction) {
                ObjectDirection.north, ObjectDirection.running_north -> North
                ObjectDirection.right, ObjectDirection.running_right -> Right
                ObjectDirection.east, ObjectDirection.running_east -> East
                ObjectDirection.down, ObjectDirection.running_down -> Down
                ObjectDirection.south, ObjectDirection.running_south -> South
                ObjectDirection.left, ObjectDirection.running_left -> Left
                ObjectDirection.west, ObjectDirection.running_west -> West
                ObjectDirection.up, ObjectDirection.running_up -> Up
                else -> null
            }
        }
    }
}