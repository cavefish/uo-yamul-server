package handlers

import (
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/commands"
)

func SendClientFeatures(client interfaces.ClientConnection, features commands.ClientFeatures) { // 0xB9
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xB9)
	client.WriteUInt(ConvertClientFeaturesToFlags(features))
}
