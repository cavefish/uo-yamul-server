package onGameLoginRequest

import (
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnLoginRequest(event listeners.CommandEvent[commands.GameLoginRequest]) {
	loginError, deniedReason := validateLogin(event.Command)
	if loginError {
		response := commands.LoginDeniedCommand{
			Reason: deniedReason,
		}
		handlers.LoginDenied(event.Client, response)
		return
	}

	ShowCharacterSelection(event.Client)
}

func validateLogin(command commands.GameLoginRequest) (bool, commands.LoginDeniedReason) {
	ok, err := login.CheckUserCredentials(command.Username, command.Password)
	if err == login.INVALID_USER {
		return ok, commands.AccountBlocked
	}
	if err == login.INVALID_CREDENTIALS {
		return ok, commands.IncorrectUsernamePassword
	}
	return ok, 0

}

func ShowCharacterSelection(client *connection.ClientConnection) {
	clientFeatures := commands.ClientFeatures{
		Unknown0001:         true,
		SingleCharacterSlot: true,
	}
	handlers.SendClientFeatures(client, clientFeatures)
	charactersStartLocation := commands.CharactersStartLocation{
		Characters:         make([]commands.CharacterLogin, 5),
		LastValidCharacter: 0,
		StartingCities:     make([]commands.StartingCity, 1),
		Flags:              clientFeatures,
	}
	for i := 0; i < len(charactersStartLocation.Characters); i++ {
		charactersStartLocation.Characters[i].Name = "asdf" //fmt.Sprintf("Username%d", i+1)
		charactersStartLocation.Characters[i].Password = ""
	}
	for i := 0; i < len(charactersStartLocation.StartingCities); i++ {
		charactersStartLocation.StartingCities[i].Name = "Yew"                //fmt.Sprintf("City%d", i+1)
		charactersStartLocation.StartingCities[i].Tavern = "The Empath Abbey" //fmt.Sprintf("Tavern%d", i+1)
		charactersStartLocation.StartingCities[i].CoordinateX = 633
		charactersStartLocation.StartingCities[i].CoordinateY = 858
		charactersStartLocation.StartingCities[i].CoordinateZ = 0
		charactersStartLocation.StartingCities[i].CoordinateMap = 0
		charactersStartLocation.StartingCities[i].ClilocDescription = 1075072

	}
	handlers.SendCharactersAndStartingLocations(client, charactersStartLocation)
}
