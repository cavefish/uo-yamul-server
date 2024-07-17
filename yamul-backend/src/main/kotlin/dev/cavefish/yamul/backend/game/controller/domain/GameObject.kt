package dev.cavefish.yamul.backend.game.controller.domain

data class GameObject(
    val id: ObjectId = 0,
    val name: String? = null,
    val parentId: ObjectId?,
    val graphicId: GraphicId,
    val layer: Int = 0,
    val hue: Hue,
    val flags: FlagsValue = 0,
    val notoriety: NotorietyValue = 0,
    val items: List<GameObjectItem> = emptyList()
)

data class GameObjectItem(
    val id: ObjectId,
    val graphicId: GraphicId,
    val hue: Hue,
    val layer: Int
)
