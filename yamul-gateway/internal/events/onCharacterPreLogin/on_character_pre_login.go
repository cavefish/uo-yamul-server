package onCharacterPreLogin

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/handlers"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnCharacterPreLogin(event listeners.CommandEvent[commands.PreLogin]) {
	event.Client.LoginDetails.CharacterSlot = event.Command.Slot
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
