package onShardSelected

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnShardSelected(event listeners.CommandEvent[commands.ShardSelected]) {
	command := commands.RedirectToShard{
		AddressIP:     event.Client.Connection.LocalAddr().String(),
		EncryptionKey: 0, //event.Client.EncryptionState.Seed,
	}
	handlers.RedirectToShard(event.Client, command)
}
