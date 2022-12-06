package handlers

import (
	"fmt"
	"reflect"
	"runtime"
	"yamul-gateway/internal/transport/multima/connection"
)

func Setup() {
	for i := 0; i < 256; i++ {
		connection.ClientCommandHandlers[i] = noop
	}
	setHandler(0x5d, preLogin)
	setHandler(0x73, ping)
	setHandler(0x80, loginRequest)
	forbiddenClientCommand(0x82, "Login denied")
	setHandler(0x91, gameServerLogin)
	setHandler(0xa0, serverSelected)
	setHandler(0xbf, genericCommand)
	setHandler(0xef, newSeed)
}

func setHandler(command byte, delegate func(client *connection.ClientConnection)) {
	handlerName := runtime.FuncForPC(reflect.ValueOf(delegate).Pointer()).Name()
	loggerPrefix := fmt.Sprintf("[%x, %s]", command, handlerName)
	handler := func(client *connection.ClientConnection, commandCode byte) {
		client.Logger.SetPrefix(loggerPrefix)
		delegate(client)
		client.Logger.SetPrefix("")
	}
	connection.ClientCommandHandlers[command] = handler
}

func unimplemented(skip int) connection.CommandHandler {
	return func(client *connection.ClientConnection, commandCode byte) {
		client.ReadFixedString(skip)
	}
}

func noop(client *connection.ClientConnection, commandCode byte) {
	client.Err = fmt.Errorf("unknown command %x", commandCode)
}

func forbiddenClientCommand(command byte, description string) {
	handler := func(client *connection.ClientConnection, commandCode byte) {
		client.Err = fmt.Errorf("forbidden command %x %s", commandCode, description)
	}
	connection.ClientCommandHandlers[command] = handler
}
