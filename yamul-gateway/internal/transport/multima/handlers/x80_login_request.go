package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/listeners"
)

func loginRequest(client *connection.ClientConnection) { // 0x80
	username := client.ReadFixedString(30)
	password := client.ReadFixedString(30)
	nextKey := client.ReadByte()
	body := commands.LoginRequestCommand{Username: username, Password: password, Nextkey: nextKey}

	listeners.Listeners.OnLoginRequest.Trigger(client, body)
}
