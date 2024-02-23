package onShardSelected

import (
	"yamul-gateway/internal/events/onGameLoginRequest"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

const redirectToGameServer = true

func OnShardSelected(event listeners.CommandEvent[commands.ShardSelected]) {
	if redirectToGameServer {
		command := commands.RedirectToShard{
			AddressIP:     event.Client.GetConnection().LocalAddr().String(),
			EncryptionKey: 0x5f11d4f0, //event.Client.encryptionState.Seed,
		}
		handlers.RedirectToShard(event.Client, command)
		return
	}

	err := onGameLoginRequest.ShowCharacterSelection(event.Client, nil)
	if err != nil {
		event.Client.KillConnection(err)
	}
}
