package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func GeneralLightLevel(client interfaces.ClientConnection, command commands.GeneralLightLevel) { // 0x4F
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x4F)
	client.WriteByte(command.Level)
}
