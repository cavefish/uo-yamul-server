package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/listeners"
)

func serverSelected(client *connection.ClientConnection) { // 0xA0
	idx := client.ReadUShort()

	body := commands.ShardSelected{Idx: idx}

	listeners.Listeners.OnShardSelected.Trigger(client, body)
}
