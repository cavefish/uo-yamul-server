package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func MoveRejectFromServer(client interfaces.ClientConnection, command commands.MoveReject) { // 0x22
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x21)
	client.WriteByte(command.Sequence)
	client.WriteUShort(command.XLoc)
	client.WriteUShort(command.YLoc)
	client.WriteByte(command.Direction)
	client.WriteByte(command.ZLoc)
}
