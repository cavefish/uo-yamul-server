package onLoginRequest

import (
	"strings"
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
		PercentageOfPlayers: 0x01,
	}
	command := commands.ListGameServers{Flags: 0xff, Servers: []commands.GameServer{server}}
	handlers.ListGameServers(event.Client, command)
}

func validateLogin(command commands.LoginRequestCommand) (bool, commands.LoginDeniedReason) {
	if !strings.EqualFold(command.Username, "admin") {
		return false, commands.AccountBlocked
	}
	if !strings.EqualFold(command.Password, "admin") {
		return false, commands.IncorrectUsernamePassword
	}
	return true, 0

}
