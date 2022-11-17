package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func SendClientFeatures(client *connection.ClientConnection, features commands.ClientFeatures) { // 0xB9
	client.Lock()
	defer client.Unlock()

	var flags uint16 = 0
	if features.Chat {
		flags |= 0x0001
	}
	if features.LbrAnimations {
		flags |= 0x0002
	}
	if features.CreatePaladinNecromancer {
		flags |= 0x0010
	}
	if features.SixthSlot {
		flags |= 0x0020
	}
	if features.ExtraFeatures {
		flags |= 0x8000
	}

	client.WriteUShort(flags)
}
