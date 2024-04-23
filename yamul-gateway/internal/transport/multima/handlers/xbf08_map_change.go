package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func MapChange(client interfaces.ClientConnection, command commands.MapChange) { // 0xbf08
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xbf)
	client.WriteUShort(6)
	client.WriteUShort(0x0008)
	client.WriteByte(command.MapId)
}
