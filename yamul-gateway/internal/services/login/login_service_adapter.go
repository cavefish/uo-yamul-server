package login

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/handlers"
)

func ValidateLogin(username string, password string) (bool, commands.LoginDeniedReason) {
	return service.CheckUserCredentials(username, password)
}

func DenyLogin(client *connection.ClientConnection, deniedReason commands.LoginDeniedReason) {
	response := commands.LoginDeniedCommand{
		Reason: deniedReason,
	}
	handlers.LoginDenied(client, response)
}
