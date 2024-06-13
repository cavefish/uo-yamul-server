package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func ServerMoveAck(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := commands.MoveAck{
		Sequence: byte(msg.Body.GetMoveAck().Sequence),
		Status:   byte(msg.Body.GetMoveAck().GetNotorietyFlags().Number()),
	}
	handlers.MoveAckFromServer(connection, command)
}
