package handlers

import (
	"yamul-gateway/internal/transport/multima/connection"
)

func LoginComplete(client *connection.ClientConnection, nothing any) { // 0x55
	client.Lock()
	defer client.Unlock()
	client.WriteByte(0x55)
}
