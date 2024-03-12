package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func MapChange(client interfaces.ClientConnection, command commands.MapChange) { // 0x6D
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(command.MapId)
}
