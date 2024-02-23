package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func LoginComplete(client interfaces.ClientConnection, nothing any) { // 0x55
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0x55)
}
