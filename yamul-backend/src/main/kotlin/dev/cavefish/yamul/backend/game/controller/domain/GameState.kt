package dev.cavefish.yamul.backend.game.controller.domain

data class GameState(
    val characterObjectId: Int = 0,
    val coordinates: Coordinates = Coordinates()
)
