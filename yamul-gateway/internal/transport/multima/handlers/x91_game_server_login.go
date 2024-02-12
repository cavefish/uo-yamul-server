package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/listeners"
)

func gameServerLogin(client *connection.ClientConnection) { // 0xA0
	encriptionKey := client.ReadUInt()
	username := client.ReadFixedString(30)
	password := client.ReadFixedString(30)
	command := commands.GameLoginRequest{
		Username:      username,
		Password:      password,
		EncryptionKey: encriptionKey,
	}

	listeners.Listeners.OnGameLoginRequest.Trigger(client, command)

}
