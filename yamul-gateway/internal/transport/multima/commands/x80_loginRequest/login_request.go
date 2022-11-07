package x80_loginRequest

import (
	"yamul-gateway/internal/transport/multima/commands/x82_loginDenied"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/listeners"
)

type LoginRequestCommand struct {
	username string
	password string
	nextkey  byte
}

func LoginRequest(client *connection.ClientConnection, commandCode byte) { // 0x80
	username := client.ReadFixedString(30)
	password := client.ReadFixedString(30)
	nextKey := client.ReadByte()
	body := LoginRequestCommand{username: username, password: password, nextkey: nextKey}
	event := listeners.Build[LoginRequestCommand](client, body)
	listeners.Trigger(OnLoginRequest, event)
}

var OnLoginRequest listeners.CommandListener[LoginRequestCommand] = onLoginRequest

func onLoginRequest(event listeners.CommandEvent[LoginRequestCommand]) {
	response := x82_loginDenied.LoginDeniedCommand{
		Reason: x82_loginDenied.IncorrectUsernamePassword,
	}
	x82_loginDenied.LoginDenied(event.Client, response)
}
