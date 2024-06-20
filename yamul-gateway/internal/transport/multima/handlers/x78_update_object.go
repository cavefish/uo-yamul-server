package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func UpdateObject(client interfaces.ClientConnection, command commands.UpdateObject) { // 0x78
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x78)
	packageSize := 23 + 9*len(command.Items)
	client.WriteUShort(uint16(packageSize)) // size
	client.WriteUInt(command.Serial)
	client.WriteUShort(command.GraphicId)
	client.WriteUShort(command.XLoc)
	client.WriteUShort(command.YLoc)
	client.WriteByte(command.ZLoc)
	client.WriteByte(command.Direction)
	client.WriteUShort(command.Hue)
	client.WriteByte(command.Flags)
	client.WriteByte(command.NotorietyFlag)
	for _, item := range command.Items {
		client.WriteUInt(item.Serial)
		client.WriteUShort(item.Artwork)
		client.WriteByte(item.Layer)
		client.WriteUShort(item.Hue)
	}
	client.WriteUInt(0) // 0 ended array of items
}
