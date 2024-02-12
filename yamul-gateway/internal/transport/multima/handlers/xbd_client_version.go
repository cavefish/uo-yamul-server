package handlers

import (
	"yamul-gateway/internal/transport/multima/connection"
)

func receiveClientVersion(client *connection.ClientConnection) { // 0xBD
	size := client.ReadUShort() - 3
	body := client.ReadFixedBytes(int(size))
	client.Logger.Debug("User logged with version %s", body)
}
