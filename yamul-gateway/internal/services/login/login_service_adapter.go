package login

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func ValidateLogin(username string, password string) (bool, commands.LoginDeniedReason) {
	return service.CheckUserCredentials(username, password)
}

func DenyLogin(client interfaces.ClientConnection, deniedReason commands.LoginDeniedReason) {
	response := commands.LoginDeniedCommand{
		Reason: deniedReason,
	}
	handlers.LoginDenied(client, response)
}
