package dev.cavefish.yamul.backend.game.controller.domain

data class GameObject(
    val id: ObjectId = 0,
    val parentId: ObjectId?,
    val graphicId: GraphicId,
    val layer: Int = 0,
    val hue: Hue,
    val flags: Flags = Flags.None,
    val notoriety: Notoriety = Notoriety.Unknown,
    val items: List<GameObjectItem> = emptyList()
)

data class GameObjectItem(
    val id: ObjectId,
    val graphicId: GraphicId,
    val hue: Hue,
    val layer: Int
)
