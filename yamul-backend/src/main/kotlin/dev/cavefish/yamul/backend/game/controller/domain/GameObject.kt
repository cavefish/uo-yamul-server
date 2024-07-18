package dev.cavefish.yamul.backend.game.controller.domain

data class GameObject(
    val id: ObjectId = 0,
    val name: String? = null,
    val isCharacter: Boolean = false,
    val parentId: ObjectId?,
    val graphicId: GraphicId,
    val layer: Int = 0,
    val hue: Hue,
    val flags: List<Flags> = emptyList(),
    val notoriety: List<Notoriety> = emptyList(),
    val items: List<GameObjectItem> = emptyList()
)

data class GameObjectItem(
    val id: ObjectId,
    val graphicId: GraphicId,
    val hue: Hue,
    val layer: Int
)
