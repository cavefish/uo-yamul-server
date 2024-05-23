package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func UpdateObject(client interfaces.ClientConnection, command commands.UpdateObject) { // 0x78
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x78)
	client.WriteUShort(23) // size
	client.WriteUInt(command.Serial)
	client.WriteUShort(command.GraphicId)
	client.WriteUShort(command.XLoc)
	client.WriteUShort(command.YLoc)
	client.WriteByte(command.ZLoc)
	client.WriteByte(command.Direction)
	client.WriteUShort(command.Hue)
	client.WriteByte(command.Flags)
	client.WriteByte(command.NotorietyFlag)
	client.WriteUInt(0) // 0 ended array of items
}
