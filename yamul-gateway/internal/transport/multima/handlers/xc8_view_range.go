package handlers

import (
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func clientViewRange(client interfaces.ClientConnection) { // 0xC8
	viewRange := client.ReadByte() // TODO implement
	listeners.Listeners.OnClientViewRange.Trigger(client, viewRange)
}
