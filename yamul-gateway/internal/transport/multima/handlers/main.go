package handlers

import (
	"fmt"
	"reflect"
	"runtime"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/connection"
)

func Setup() {
	for i := 0; i < 256; i++ {
		connection.ClientCommandHandlers[i] = noop
	}
	setHandler(0x09, unimplemented(4)) // Single click on event id http://www.hoogi.de/wolfpack/wiki/doku.php?id=uo_protocol_0x09
	setHandler(0x34, unimplemented(9)) // Get player status http://www.hoogi.de/wolfpack/wiki/doku.php?id=uo_protocol_0x34
	setHandler(0x5d, preLogin)
	setHandler(0x73, ping)
	setHandler(0x80, loginRequest)
	forbiddenClientCommand(0x82, "Login denied")
	setHandler(0x91, gameServerLogin)
	setHandler(0xa0, serverSelected)
	setHandler(0xbd, receiveClientVersion)
	setHandler(0xbf, receiveGenericCommand)
	setHandler(0xef, newSeed)
	setHandler(0xfb, useMultiSight)
}

func setHandler(command byte, delegate func(client interfaces.ClientConnection)) {
	handlerName := runtime.FuncForPC(reflect.ValueOf(delegate).Pointer()).Name()
	commandAsText := fmt.Sprintf("0x%02x", command)
	handler := func(client interfaces.ClientConnection, commandCode byte) {
		logger := client.GetLogger()
		logger.SetLogField("command", commandAsText)
		defer logger.ClearLogField("command")
		logger.SetLogField("handler", handlerName)
		defer logger.ClearLogField("handler")
		delegate(client)
	}
	connection.ClientCommandHandlers[command] = handler
}

func unimplemented(skip int) func(client interfaces.ClientConnection) {
	return func(client interfaces.ClientConnection) {
		client.GetLogger().Infof("Unimplemented command, skipping %d bytes.", skip)
		client.ReadFixedString(skip)
	}
}

func noop(client interfaces.ClientConnection, commandCode byte) {
	client.KillConnection(fmt.Errorf("unknown command %x", commandCode))
}

func forbiddenClientCommand(command byte, description string) {
	handler := func(client interfaces.ClientConnection, commandCode byte) {
		client.KillConnection(fmt.Errorf("forbidden command %x %s", commandCode, description))
	}
	connection.ClientCommandHandlers[command] = handler
}
