package clientEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnUnicodeSpeechRequest(event listeners.CommandEvent[commands.UnicodeSpeechSelected]) {
	keywords := make([]uint32, len(event.Command.Keywords))
	for i := 0; i < len(keywords); i++ {
		keywords[i] = uint32(event.Command.Keywords[i])
	}
	request := services.Message_UnicodeSpeechSelected{
		UnicodeSpeechSelected: &services.MsgUnicodeSpeechSelected{
			Mode:     services.MessageType(event.Command.Mode),
			Hue:      uint32(event.Command.Hue),
			Font:     services.Fonts(event.Command.Font),
			Language: event.Command.Language,
			Keywords: keywords,
			Text:     event.Command.Text,
		},
	}

	event.Client.GetGameService().Send(services.MsgType_TypeUnicodeSpeechSelected, &services.Message{Msg: &request})
}
