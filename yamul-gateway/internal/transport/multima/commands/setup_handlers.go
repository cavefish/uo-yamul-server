package commands

import (
	"fmt"
	"yamul-gateway/internal/transport/multima/commands/x80_loginRequest"
	"yamul-gateway/internal/transport/multima/commands/xef_newSeed"
	"yamul-gateway/internal/transport/multima/connection"
)

func SetupCommandHandlers() {
	for i := 0; i < 256; i++ {
		connection.ClientCommandHandlers[i] = noop
	}
	connection.ClientCommandHandlers[0x80] = x80_loginRequest.LoginRequest
	connection.ClientCommandHandlers[0x82] = forbiddenClientCommand
	connection.ClientCommandHandlers[0xef] = xef_newSeed.NewSeed
}

func noop(client *connection.ClientConnection, commandCode byte) {
	client.Err = fmt.Errorf("unknown command %x", commandCode)
}

func forbiddenClientCommand(client *connection.ClientConnection, commandCode byte) {
	client.Err = fmt.Errorf("forbidden command %x", commandCode)
}
