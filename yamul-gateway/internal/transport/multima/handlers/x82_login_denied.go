package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func LoginDenied(client *connection.ClientConnection, response commands.LoginDeniedCommand) { // 0x82
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0x82)
	client.WriteByte(byte(response.Reason))
	_ = client.SendAnyData()
	client.ShouldCloseConnection = true
}
