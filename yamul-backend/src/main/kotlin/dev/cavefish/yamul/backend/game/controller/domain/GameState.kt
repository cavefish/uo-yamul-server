package dev.cavefish.yamul.backend.game.controller.domain

data class GameState(
    val characterObject: GameObject,
    val coordinates: Coordinates = Coordinates()
)
