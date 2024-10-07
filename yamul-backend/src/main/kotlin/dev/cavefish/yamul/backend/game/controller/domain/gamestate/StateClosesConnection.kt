package dev.cavefish.yamul.backend.game.controller.domain.gamestate


sealed interface StateClosesConnection

data object StateLoggedOut : StateClosesConnection, State(
    properties = listOf()
)

open class StateError(val errorDescription: String) : StateClosesConnection, State(
    properties = listOf()
)

data object StateErrorRequiresLoggedIn : StateError("Required LoggedIn state")
data object StateErrorRequiresCharacter : StateError("Required StateHasCharacter state")
