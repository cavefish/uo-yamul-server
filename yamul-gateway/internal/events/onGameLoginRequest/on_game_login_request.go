package onGameLoginRequest

import (
	"fmt"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/commands"
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

	clientFeatures := commands.ClientFeatures{}
	handlers.SendClientFeatures(event.Client, clientFeatures)
	charactersStartLocation := commands.CharactersStartLocation{
		Characters:     make([]commands.CharacterLogin, 5),
		StartingCities: make([]commands.StartingCity, 12),
		Flags:          clientFeatures,
	}
	for i := 0; i < len(charactersStartLocation.Characters); i++ {
		charactersStartLocation.Characters[i].Name = fmt.Sprintf("Username #%d", i+1)
		charactersStartLocation.Characters[i].Password = ""
	}
	handlers.SendCharactersAndStartingLocations(event.Client, charactersStartLocation)
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
