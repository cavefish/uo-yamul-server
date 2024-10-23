package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func ServerMoveReject(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	body := msg.Body.GetMoveReject()
	command := commands.MoveReject{
		Sequence:  byte(body.Sequence),
		XLoc:      uint16(body.XLoc),
		YLoc:      uint16(body.YLoc),
		ZLoc:      byte(body.ZLoc),
		Direction: byte(body.Direction.Number()),
	}
	handlers.MoveRejectFromServer(connection, command)
}
