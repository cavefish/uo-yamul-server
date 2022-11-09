package handlers

import (
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/listeners"
	"yamul-gateway/internal/transport/multima/messages"
)

func LoginRequest(client *connection.ClientConnection, commandCode byte) { // 0x80
	username := client.ReadFixedString(30)
	password := client.ReadFixedString(30)
	nextKey := client.ReadByte()
	body := messages.LoginRequestCommand{Username: username, Password: password, Nextkey: nextKey}
	event := listeners.Build[messages.LoginRequestCommand](client, body)
	listeners.Trigger(OnLoginRequest, event)
}

var OnLoginRequest listeners.CommandListener[messages.LoginRequestCommand] = onLoginRequest

func onLoginRequest(event listeners.CommandEvent[messages.LoginRequestCommand]) {
	response := messages.LoginDeniedCommand{
		Reason: messages.IncorrectUsernamePassword,
	}
	LoginDenied(event.Client, response)
}
