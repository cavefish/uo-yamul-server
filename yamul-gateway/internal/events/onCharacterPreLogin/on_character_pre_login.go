package onCharacterPreLogin

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnCharacterPreLogin(event listeners.CommandEvent[commands.PreLogin]) {
	success, deniedReason := login.ValidateLogin(event.Client.GetLoginDetails().Username, event.Command.Password)
	if !success {
		login.DenyLogin(event.Client, deniedReason)
		return
	}

	event.Client.GetLoginDetails().CharacterSlot = event.Command.Slot

	err := event.Client.CreateGameConnection()
	if err != nil {
		event.Client.KillConnection(err)
		return
	}

	command := commands.PlayerStartConfirmation{
		CharacterID:       1,
		CharacterBodyType: 0x0190,
		Coordinates: commands.Coordinates{
			X: 0x041D,
			Y: 0x0598,
			Z: 0xFFAB,
		},
		DirectionFacing: commands.DirectionFacing{
			Direction: 0x07,
		},
	}
	handlers.PlayerStartConfirmation(event.Client, command)
}
