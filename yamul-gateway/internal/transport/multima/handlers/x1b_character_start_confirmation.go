package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func PlayerStartConfirmation(client interfaces.ClientConnection, body commands.PlayerStartConfirmation) { // 0x1B
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x1B)
	client.WriteUInt(body.CharacterID)
	client.WriteUInt(0)
	client.WriteUShort(body.CharacterBodyType)
	client.WriteUShort(body.Coordinates.X)
	client.WriteUShort(body.Coordinates.Y)
	client.WriteUShort(body.Coordinates.Z)
	client.WriteByte(body.DirectionFacing.Direction)
	client.WriteByte(0)
	client.WriteUInt(0xFFFFFFFF)
	client.WriteUShort(0)
	client.WriteUShort(0)
	client.WriteUShort(0x1800)
	client.WriteUShort(0x1000)
	client.WriteUShort(0)
	client.WriteUInt(0)
}
