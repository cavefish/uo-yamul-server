package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnPlayerStartConfirmation(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := toPlayerStartConfirmationCommand(msg.Body.GetPlayerStartConfirmation())
	handlers.PlayerStartConfirmation(connection, command)
}

func toPlayerStartConfirmationCommand(confirmation *services.MsgPlayerStartConfirmation) commands.PlayerStartConfirmation {
	return commands.PlayerStartConfirmation{
		CharacterID:       confirmation.Id.Value,
		CharacterBodyType: 0,
		Coordinates: commands.Coordinates{
			X: uint16(confirmation.Coordinates.XLoc),
			Y: uint16(confirmation.Coordinates.YLoc),
			Z: uint16(confirmation.Coordinates.ZLoc),
		},
		DirectionFacing: commands.DirectionFacing{
			Direction: byte(confirmation.Direction.Number()),
		},
	}
}
