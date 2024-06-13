package clientEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnClientMoveRequest(event listeners.CommandEvent[commands.ClientMoveRequest]) {

	selection := services.Message_ClientMoveRequest{
		ClientMoveRequest: &services.MsgClientMoveRequest{
			Direction: services.ObjectDirection(event.Command.Direction),
			Sequence:  uint32(event.Command.Sequence),
			AckKey:    event.Command.AckKey,
		},
	}
	event.Client.GetGameService().Send(services.MsgType_TypeClientMoveRequest, &services.Message{Msg: &selection})
}
