package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnMapChange(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	var command = toCommandMapChange(msg.Body.GetMapChange())
	handlers.MapChange(connection, command)
}

func toCommandMapChange(body *services.MsgMapChange) commands.MapChange {
	return commands.MapChange{
		MapId: byte(body.MapId),
	}
}
