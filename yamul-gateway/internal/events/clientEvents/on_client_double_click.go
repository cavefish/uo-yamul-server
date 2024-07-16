package clientEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnClientDoubleClick(event listeners.CommandEvent[uint32]) {

	msg := services.Message_ClientDoubleClick{
		ClientDoubleClick: &services.MsgClientDoubleClick{
			Target: &services.ObjectId{
				Value: event.Command,
			},
		},
	}

	event.Client.GetGameService().Send(services.MsgType_TypeClientDoubleClick, &services.Message{Msg: &msg})
}
