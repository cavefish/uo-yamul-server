package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnUpdateObject(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := toCommandUpdateObject(msg.Body.GetUpdateObject())
	handlers.UpdateObject(connection, command)
}

func toCommandUpdateObject(msg *services.MsgUpdateObject) commands.UpdateObject {
	return commands.UpdateObject{
		Serial:        msg.Id.Value,
		GraphicId:     uint16(msg.GraphicId),
		XLoc:          uint16(msg.XLoc),
		YLoc:          uint16(msg.YLoc),
		ZLoc:          byte(msg.ZLoc),
		Direction:     byte(msg.Direction),
		Hue:           uint16(msg.Hue),
		Flags:         byte(msg.Flags),
		NotorietyFlag: byte(msg.NotorietyFlags),
	}
}
