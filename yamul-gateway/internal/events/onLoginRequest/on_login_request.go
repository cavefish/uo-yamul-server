package onLoginRequest

import (
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnLoginRequest(event listeners.CommandEvent[commands.LoginRequestCommand]) {
	loginError, deniedReason := validateLogin(event.Command)
	if loginError {
		response := commands.LoginDeniedCommand{
			Reason: deniedReason,
		}
		handlers.LoginDenied(event.Client, response)
		return
	}

	server := commands.GameServer{
		Name:                "This server",
		AddressIP:           event.Client.Connection.LocalAddr().String(),
		Timezone:            0x00,
		PercentageOfPlayers: 0x00,
	}
	command := commands.ListGameServers{Flags: 0xff, Servers: []commands.GameServer{server}}
	handlers.ListGameServers(event.Client, command)
}

func validateLogin(command commands.LoginRequestCommand) (bool, commands.LoginDeniedReason) {
	ok, err := login.CheckUserCredentials(command.Username, command.Password)
	if err == login.INVALID_USER {
		return ok, commands.AccountBlocked
	}
	if err == login.INVALID_CREDENTIALS {
		return ok, commands.IncorrectUsernamePassword
	}
	return ok, 0

}
