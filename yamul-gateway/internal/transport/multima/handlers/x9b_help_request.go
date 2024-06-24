package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func helpRequest(client interfaces.ClientConnection) { // 0x9B
	_ = client.ReadFixedBytes(257) // Ignored payload
	client.GetLogger().Info("Unimplemented \"Help request\" handler")
}
