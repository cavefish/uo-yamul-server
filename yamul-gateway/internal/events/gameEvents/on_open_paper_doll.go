package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnOpenPaperDoll(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := mapToOpenPaperDoll(msg.Body.GetOpenPaperDoll())
	handlers.OpenPaperDoll(connection, command)
}

func mapToOpenPaperDoll(input *services.MsgOpenPaperDoll) commands.OpenPaperDoll {
	var f byte = 0
	for _, flag := range input.Flags {
		f = f | byte(flag)
	}
	return commands.OpenPaperDoll{
		Id:     input.Id.Value,
		Name:   input.Name,
		Status: f,
	}
}
