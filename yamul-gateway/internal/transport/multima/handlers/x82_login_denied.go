package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func LoginDenied(client interfaces.ClientConnection, response commands.LoginDeniedCommand) { // 0x82
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0x82)
	client.WriteByte(byte(response.Reason))
	_ = client.SendAnyData()
	client.KillConnection(nil)
}
