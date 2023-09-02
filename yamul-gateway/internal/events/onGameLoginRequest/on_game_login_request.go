package onGameLoginRequest

import (
	"yamul-gateway/internal/services/character"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnLoginRequest(event listeners.CommandEvent[commands.GameLoginRequest]) {
	success, deniedReason := validateLogin(event.Command)
	if !success {
		denyLogin(event, deniedReason)
		return
	}
	event.Client.SetLogin(event.Command.Username, event.Command.Password)

	service, err := character.NewCharacterService(event.Client)
	if err != nil {
		denyLogin(event, commands.LoginDeniedReason_CommunicationProblem)
		return
	}
	defer service.Close()

	err = ShowCharacterSelection(event.Client, service)
	if err != nil {
		denyLogin(event, commands.LoginDeniedReason_CommunicationProblem)
		return
	}

	return

}

func denyLogin(event listeners.CommandEvent[commands.GameLoginRequest], deniedReason commands.LoginDeniedReason) {
	response := commands.LoginDeniedCommand{
		Reason: deniedReason,
	}
	handlers.LoginDenied(event.Client, response)
}

func validateLogin(command commands.GameLoginRequest) (bool, commands.LoginDeniedReason) {
	return login.Service.CheckUserCredentials(command.Username, command.Password)
}

func ShowCharacterSelection(client *connection.ClientConnection, service *character.CharacterService) error {
	clientFeatures := commands.ClientFeatures{
		Unknown0001:         true,
		SingleCharacterSlot: true,
	}
	handlers.SendClientFeatures(client, clientFeatures)
	characters, lastValidCharacter, err := service.GetCharacters()
	if err != nil {
		return err
	}
	charactersStartLocation := commands.CharactersStartLocation{
		Characters:         characters,
		LastValidCharacter: lastValidCharacter,
		StartingCities:     make([]commands.StartingCity, 1),
		Flags:              clientFeatures,
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
	return nil
}
