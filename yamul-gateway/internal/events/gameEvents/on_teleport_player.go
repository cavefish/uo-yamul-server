package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnTeleportPlayer(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := toCommandTeleportPlayer(msg.Body.GetTeleportPlayer())
	handlers.TeleportPlayer(connection, command)
}

func toCommandTeleportPlayer(player *services.MsgTeleportPlayer) commands.TeleportPlayer {
	var status byte = 0
	for i := range player.Status {
		status = status | byte(player.Status[i].Number())
	}
	return commands.TeleportPlayer{
		Serial:    player.Id.Value,
		Status:    status,
		XLoc:      uint16(player.Coordinates.XLoc),
		YLoc:      uint16(player.Coordinates.YLoc),
		Direction: byte(player.Direction.Number()),
		ZLoc:      int8(player.Coordinates.ZLoc),
	}
}
