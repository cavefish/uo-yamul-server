package clientEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnClientMoveAck(event listeners.CommandEvent[commands.MoveAck]) {
	event.Client.GetLogger().Infof("Received move ack from client: %v", event.Command)
	selection := services.Message_MoveAck{
		MoveAck: &services.MsgMoveAck{
			Sequence:       uint32(event.Command.Sequence),
			NotorietyFlags: services.Notoriety(event.Command.Status),
		},
	}
	event.Client.GetGameService().Send(services.MsgType_TypeClientMoveRequest, &services.Message{Msg: &selection})
}
