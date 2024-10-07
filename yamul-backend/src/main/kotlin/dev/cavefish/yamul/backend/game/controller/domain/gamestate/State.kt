package dev.cavefish.yamul.backend.game.controller.domain.gamestate

sealed class State(private val properties: List<StateProperty>) {
    open fun hasProperty(property: StateProperty) = properties.contains(property)
    fun hasWonProperty(previousState: State, property: StateProperty): Boolean {
        if (!hasProperty(property)) {
            return false
        }
        if (previousState.hasProperty(property)) {
            return false
        }
        return true
    }

    fun hasLostProperty(previousState: State, property: StateProperty): Boolean {
        if (hasProperty(property)) {
            return false
        }
        if (!previousState.hasProperty(property)) {
            return false
        }
        return true
    }
}