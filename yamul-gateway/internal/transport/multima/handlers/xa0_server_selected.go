package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func serverSelected(client interfaces.ClientConnection) { // 0xA0
	idx := client.ReadUShort()

	body := commands.ShardSelected{Idx: idx}

	listeners.Listeners.OnShardSelected.Trigger(client, body)
}
