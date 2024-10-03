package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func SystemSendText(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.SystemSendText(connection, mapToCommandSystemSendText(msg.Body.GetSystemSendText()))
}

func mapToCommandSystemSendText(body *services.MsgSystemSendText) commands.SystemSendText {
	serial := uint32(0xFFFFFFFF)
	model := uint16(0xFFFF)
	if body.Id != nil {
		serial = body.Id.Value
		model = uint16(body.Model)
	}
	return commands.SystemSendText{
		Serial: serial,
		Model:  model,
		Type:   byte(body.Type),
		Hue:    uint16(body.Hue),
		Font:   uint16(body.Font),
		Name:   body.Name,
		Body:   body.Body,
	}
}
