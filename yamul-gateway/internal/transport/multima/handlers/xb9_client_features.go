package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func SendClientFeatures(client *connection.ClientConnection, features commands.ClientFeatures) { // 0xB9
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xB9)
	client.WriteUInt(ConvertClientFeaturesToFlags(features))
}
