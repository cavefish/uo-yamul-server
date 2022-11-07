package x80_loginRequest

import (
	"yamul-gateway/internal/transport/multima/commands/x82_loginDenied"
	"yamul-gateway/internal/transport/multima/connection"
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
	go OnLoginRequest(client, body)
}

func OnLoginRequest(client *connection.ClientConnection, command LoginRequestCommand) {
	response := x82_loginDenied.LoginDeniedCommand{
		Reason: x82_loginDenied.CommunicationProblem,
	}
	x82_loginDenied.LoginDenied(client, response)
}
