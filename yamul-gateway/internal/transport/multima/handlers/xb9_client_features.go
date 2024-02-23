package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func SendClientFeatures(client interfaces.ClientConnection, features commands.ClientFeatures) { // 0xB9
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xB9)
	client.WriteUInt(ConvertClientFeaturesToFlags(features))
}
