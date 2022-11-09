package onLoginRequest

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnLoginRequest(event listeners.CommandEvent[commands.LoginRequestCommand]) {
	response := commands.LoginDeniedCommand{
		Reason: commands.IncorrectUsernamePassword,
	}
	handlers.LoginDenied(event.Client, response)
}
