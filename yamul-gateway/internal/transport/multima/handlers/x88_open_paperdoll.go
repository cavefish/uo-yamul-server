package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func OpenPaperDoll(client interfaces.ClientConnection, response commands.OpenPaperDoll) { // 0x88
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x88)
	client.WriteUInt(response.Id)
	client.WriteFixedString(60, response.Name)
	client.WriteByte(response.Status)
}
