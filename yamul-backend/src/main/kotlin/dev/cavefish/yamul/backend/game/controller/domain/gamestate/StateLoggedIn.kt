package dev.cavefish.yamul.backend.game.controller.domain.gamestate

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser


sealed class StateLoggedIn(private val loggedUser: LoggedUser, properties: List<StateProperty>) : State(properties) {
    fun getLoggedUser() = loggedUser
    abstract fun assignCharacter(
        characterObject: GameObject,
        coordinates: Coordinates
    ): StateHasCharacter
}

data class StateInitial(private val loggedUser: LoggedUser) : StateLoggedIn(
    loggedUser,
    properties = listOf()
) {
    override fun assignCharacter(
        characterObject: GameObject,
        coordinates: Coordinates
    ) = StateHasCharacter(loggedUser, characterObject, coordinates)
}

data class StateHasCharacter(
    private val loggedUser: LoggedUser,
    val characterObject: GameObject,
    val coordinates: Coordinates
) : StateLoggedIn(
    loggedUser,
    properties = listOf(StateProperty.RECEIVE_EVENTS)
) {
    override fun assignCharacter(characterObject: GameObject, coordinates: Coordinates): StateHasCharacter = this.copy(
        characterObject = characterObject,
        coordinates = coordinates
    )
}
