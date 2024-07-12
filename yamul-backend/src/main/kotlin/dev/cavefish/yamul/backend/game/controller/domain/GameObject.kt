package dev.cavefish.yamul.backend.game.controller.domain

data class GameObject(
    val id: ObjectId,
    val graphicId: GraphicId,
    val hue: Hue,
    val coordinates: Coordinates,
    val flags: Flags,
    val notoriety: Notoriety,
    val items: List<GameObjectItem>
)

data class GameObjectItem(
    val id: ObjectId,
    val graphicId: GraphicId,
    val hue: Hue,
    val layer: Int
)
