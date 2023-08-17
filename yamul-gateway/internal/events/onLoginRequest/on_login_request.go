package onLoginRequest

import (
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnLoginRequest(event listeners.CommandEvent[commands.LoginRequestCommand]) {
	loginSuccess, deniedReason := validateLogin(event.Command)
	if !loginSuccess {
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
	return login.CheckUserCredentials(command.Username, command.Password)
}
