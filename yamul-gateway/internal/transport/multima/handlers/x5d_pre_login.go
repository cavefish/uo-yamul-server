package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
	"yamul-gateway/utils/stringUtils"
)

func preLogin(client interfaces.ClientConnection) { // 0x5d
	command := preLoginReadBuffer(client)

	listeners.Listeners.OnPreLogin.Trigger(client, command)
}

func preLoginReadBuffer(client interfaces.ClientConnection) commands.PreLogin {
	_ = client.ReadUInt()
	charName := stringUtils.TrimRight(client.ReadFixedString(30))
	charPassword := stringUtils.TrimRight(client.ReadFixedString(30))
	slot := client.ReadUInt()
	encryptionKey := client.ReadUInt()

	command := commands.PreLogin{
		Name:          charName,
		Password:      charPassword,
		Slot:          slot,
		EncryptionKey: encryptionKey,
	}
	return command
}
