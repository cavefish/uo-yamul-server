package dev.cavefish.yamul.backend.game.controller.domain

data class GameObjectTree(val value: GameObject, val children: List<GameObjectTree>) {
    companion object {
        fun treeOf(gameObject: GameObject) = GameObjectTree(gameObject, listOf())
        fun treeOf(gameObject: GameObject, children: List<GameObjectTree>) = GameObjectTree(gameObject, children)
        fun treeOf(relation: Pair<GameObject, List<GameObjectTree>>) = GameObjectTree(relation.first, relation.second)
    }
}