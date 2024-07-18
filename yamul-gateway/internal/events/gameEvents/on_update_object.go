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
	var f byte = 0
	for _, flag := range msg.Flags {
		f = f | byte(flag)
	}
	var n byte = 0
	for _, flag := range msg.NotorietyFlags {
		n = n | byte(flag)
	}
	command := commands.UpdateObject{
		Serial:        msg.Id.Value,
		GraphicId:     uint16(msg.GraphicId),
		XLoc:          uint16(msg.XLoc),
		YLoc:          uint16(msg.YLoc),
		ZLoc:          byte(msg.ZLoc),
		Direction:     byte(msg.Direction),
		Hue:           uint16(msg.Hue),
		Flags:         f,
		NotorietyFlag: n,
		Items:         make([]commands.UpdateObjectItem, len(msg.Items)),
	}
	for i, item := range msg.Items {
		command.Items[i] = commands.UpdateObjectItem{
			Serial:  item.Id.Value,
			Artwork: uint16(item.GraphicId),
			Layer:   byte(item.Layer),
			Hue:     uint16(item.Hue),
		}
	}
	return command
}
