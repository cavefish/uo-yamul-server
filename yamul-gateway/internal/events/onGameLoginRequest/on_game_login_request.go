package onGameLoginRequest

import (
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnLoginRequest(event listeners.CommandEvent[commands.GameLoginRequest]) {
	loginError, deniedReason := validateLogin(event.Command)
	if loginError {
		response := commands.LoginDeniedCommand{
			Reason: deniedReason,
		}
		handlers.LoginDenied(event.Client, response)
		return
	}

	command := commands.ClientFeatures{}
	handlers.SendClientFeatures(event.Client, command)
}

func validateLogin(command commands.GameLoginRequest) (bool, commands.LoginDeniedReason) {
	ok, err := login.CheckUserCredentials(command.Username, command.Password)
	if err == login.INVALID_USER {
		return ok, commands.AccountBlocked
	}
	if err == login.INVALID_CREDENTIALS {
		return ok, commands.IncorrectUsernamePassword
	}
	return ok, 0

}
