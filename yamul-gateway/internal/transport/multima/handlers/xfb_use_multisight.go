package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func useMultiSight(client interfaces.ClientConnection) { // 0xFB
	value := client.ReadByte()
	client.GetStatus().UseMultiSight = value > 0
}
