package handlers

import (
	"yamul-gateway/internal/transport/multima/connection"
)

func useMultiSight(client *connection.ClientConnection) { // 0xBD
	value := client.ReadByte()
	client.Status.UseMultiSight = value > 0
}
