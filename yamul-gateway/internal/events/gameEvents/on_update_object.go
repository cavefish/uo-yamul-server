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

func toCommandUpdateObject(object *services.MsgUpdateObject) commands.UpdateObject {
	return commands.UpdateObject{
		Serial:        object.Id.Value,
		GraphicId:     0,
		XLoc:          0,
		YLoc:          0,
		ZLoc:          0,
		Direction:     0,
		Hue:           0,
		Flags:         0,
		NotorietyFlag: 0,
	}
}
