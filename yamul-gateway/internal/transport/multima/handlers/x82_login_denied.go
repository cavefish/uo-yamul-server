package handlers

import (
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/commands"
)

func LoginDenied(client interfaces.ClientConnection, response commands.LoginDeniedCommand) { // 0x82
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0x82)
	client.WriteByte(byte(response.Reason))
	_ = client.SendAnyData()
	client.KillConnection(nil)
}
