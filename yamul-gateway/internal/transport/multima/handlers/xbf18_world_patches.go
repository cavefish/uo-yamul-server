package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func WorldPatches(client interfaces.ClientConnection, command commands.WorldPatches) { // 0xbf18
	// TODO implement something if needed
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xbf)
	client.WriteUShort(49)
	client.WriteUShort(0x0018)
	client.WriteUInt(5)
	for i := 0; i < 40; i++ {
		client.WriteByte(0)
	}
}
