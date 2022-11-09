package onShardSelected

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnShardSelected(event listeners.CommandEvent[commands.ShardSelected]) {
	command := commands.RedirectToShard{
		AddressIP:     event.Client.Connection.LocalAddr().String(),
		EncryptionKey: event.Client.EncryptSeed.Seed,
	}
	handlers.RedirectToShard(event.Client, command)
}
