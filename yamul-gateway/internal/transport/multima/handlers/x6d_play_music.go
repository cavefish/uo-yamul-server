package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func PlayMusic(client interfaces.ClientConnection, command commands.PlayMusic) { // 0x6D
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x6D)
	client.WriteUShort(command.MusicId)
}
