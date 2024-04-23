package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func TeleportPlayer(client interfaces.ClientConnection, command commands.TeleportPlayer) { // 0x20
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x20)
	client.WriteUInt(command.Serial)
	client.WriteUShort(0)  // body type
	client.WriteByte(0x00) // Unknown
	client.WriteUShort(0)  //skin color
	client.WriteByte(command.Status)
	client.WriteUShort(command.XLoc)
	client.WriteUShort(command.YLoc)
	client.WriteByte(0x00) // UNKNOWN
	client.WriteByte(0x00) // UNKNOWN
	client.WriteByte(command.Direction)
	client.WriteByte(byte(command.ZLoc))
}
