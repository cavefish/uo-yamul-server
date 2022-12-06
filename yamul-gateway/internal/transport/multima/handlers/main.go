package handlers

import (
	"fmt"
	"yamul-gateway/internal/transport/multima/connection"
)

func Setup() {
	for i := 0; i < 256; i++ {
		connection.ClientCommandHandlers[i] = noop
	}
	connection.ClientCommandHandlers[0x5d] = wrap(preLogin)
	connection.ClientCommandHandlers[0x73] = wrap(ping)
	connection.ClientCommandHandlers[0x80] = wrap(loginRequest)
	connection.ClientCommandHandlers[0x82] = forbiddenClientCommand("Login denied")
	connection.ClientCommandHandlers[0x91] = wrap(gameServerLogin)
	connection.ClientCommandHandlers[0xa0] = wrap(serverSelected)
	connection.ClientCommandHandlers[0xef] = wrap(newSeed)
}

func wrap(delegate func(client *connection.ClientConnection)) connection.CommandHandler {
	return func(client *connection.ClientConnection, commandCode byte) {
		client.Logger.Debug("> Processing command %x", commandCode)
		delegate(client)
		client.Logger.Debug("< Finished command %x", commandCode)
	}
}

func unimplemented(skip int) connection.CommandHandler {
	return func(client *connection.ClientConnection, commandCode byte) {
		client.ReadFixedString(skip)
	}
}

func noop(client *connection.ClientConnection, commandCode byte) {
	client.Err = fmt.Errorf("unknown command %x", commandCode)
}

func forbiddenClientCommand(description string) connection.CommandHandler {
	return func(client *connection.ClientConnection, commandCode byte) {
		client.Err = fmt.Errorf("forbidden command %x %s", commandCode, description)
	}
}
