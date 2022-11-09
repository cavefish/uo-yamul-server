package handlers

import (
	"fmt"
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

func Setup() {
	for i := 0; i < 256; i++ {
		connection.ClientCommandHandlers[i] = noop
	}
	connection.ClientCommandHandlers[0x80] = wrap(loginRequest)
	connection.ClientCommandHandlers[0x82] = forbiddenClientCommand("Login denied")
	connection.ClientCommandHandlers[0xef] = wrap(newSeed)
}

func wrap(delegate func(client *connection.ClientConnection)) connection.CommandHandler {
	return func(client *connection.ClientConnection, commandCode byte) {
		logging.Debug("%x ", commandCode)
		delegate(client)
		logging.Debug("\n")
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
