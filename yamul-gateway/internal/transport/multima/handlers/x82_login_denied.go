package handlers

import (
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/messages"
)

func LoginDenied(client *connection.ClientConnection, response messages.LoginDeniedCommand) { // 0x82
	client.Lock()
	defer client.Unlock()
	client.WriteByte(0x82)
	client.WriteByte(byte(response.Reason))
	_ = client.SendAnyData()
	client.ShouldCloseConnection = true
}
