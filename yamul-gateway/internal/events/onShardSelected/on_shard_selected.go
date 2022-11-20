package onShardSelected

import (
	"fmt"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

const redirectToGameServer = true

func OnShardSelected(event listeners.CommandEvent[commands.ShardSelected]) {
	if redirectToGameServer {
		command := commands.RedirectToShard{
			AddressIP:     event.Client.Connection.LocalAddr().String(),
			EncryptionKey: event.Client.EncryptionState.Seed,
		}
		handlers.RedirectToShard(event.Client, command)
		return
	}

	clientFeatures := commands.ClientFeatures{
		SingleCharacterSlot: true,
		ContextMenus:        true,
		EnableAOS:           true,
	}
	handlers.SendClientFeatures(event.Client, clientFeatures)
	charactersStartLocation := commands.CharactersStartLocation{
		Characters:         make([]commands.CharacterLogin, 5),
		LastValidCharacter: 0,
		StartingCities:     make([]commands.StartingCity, 0x12),
		Flags:              clientFeatures,
	}
	for i := 0; i < len(charactersStartLocation.Characters); i++ {
		charactersStartLocation.Characters[i].Name = "asdf" //fmt.Sprintf("Username #%d", i+1)
		charactersStartLocation.Characters[i].Password = ""
	}
	for i := 0; i < len(charactersStartLocation.StartingCities); i++ {
		charactersStartLocation.StartingCities[i].Name = fmt.Sprintf("City #%d", i+1)
		charactersStartLocation.StartingCities[i].Tavern = fmt.Sprintf("Tavern #%d", i+1)
		charactersStartLocation.StartingCities[i].CoordinateX = 1496
		charactersStartLocation.StartingCities[i].CoordinateY = 1628
		charactersStartLocation.StartingCities[i].CoordinateZ = 10
		charactersStartLocation.StartingCities[i].CoordinateMap = 0
		charactersStartLocation.StartingCities[i].ClilocDescription = 1075074

	}
	handlers.SendCharactersAndStartingLocations(event.Client, charactersStartLocation)
}
