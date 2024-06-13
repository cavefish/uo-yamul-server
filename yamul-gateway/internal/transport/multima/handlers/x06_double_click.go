package handlers

import (
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func clientDoubleClick(client interfaces.ClientConnection) { // 0x06
	target := client.ReadUInt() // TODO implement
	listeners.OnClientDoubleClick.Trigger(client, target)
}
