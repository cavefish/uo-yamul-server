package commands

import (
	"yamul-gateway/internal/transport/multima/connection"
)

func OnLoginRequest(client *connection.ClientConnection, command LoginRequestCommand) {
	response := LoginDeniedCommand{
		reason: communicationProblem,
	}
	loginDenied(client, response)
}
