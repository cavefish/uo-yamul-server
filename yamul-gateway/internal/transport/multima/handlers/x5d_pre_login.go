package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func preLogin(client interfaces.ClientConnection) { // 0x5d
	_ = client.ReadUInt()
	charName := client.ReadFixedString(30)
	charPassword := client.ReadFixedString(30)
	slot := client.ReadUInt()
	encryptionKey := client.ReadUInt()

	command := commands.PreLogin{
		Name:          charName,
		Password:      charPassword,
		Slot:          slot,
		EncryptionKey: encryptionKey,
	}

	listeners.Listeners.OnPreLogin.Trigger(client, command)
}
