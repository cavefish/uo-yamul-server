package dev.cavefish.yamul.backend.game.controller.domain

data class GameState(
    val characterObjectId: Int = 0,
    val characterBodyType: GraphicIds = GraphicIds.BodyHumanMale,
    val characterBodyHue: Hue = Hues.Red.hue,
    val coordinates: Coordinates = Coordinates()
)
