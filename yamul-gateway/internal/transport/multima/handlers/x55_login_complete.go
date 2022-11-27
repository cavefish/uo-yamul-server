package handlers

import (
	"yamul-gateway/internal/transport/multima/connection"
)

func LoginComplete(client *connection.ClientConnection, nothing any) { // 0x55
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0x55)
}
