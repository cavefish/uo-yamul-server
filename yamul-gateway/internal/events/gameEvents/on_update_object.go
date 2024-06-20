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
	command := commands.UpdateObject{
		Serial:        msg.Id.Value,
		GraphicId:     uint16(msg.GraphicId),
		XLoc:          uint16(msg.XLoc),
		YLoc:          uint16(msg.YLoc),
		ZLoc:          byte(msg.ZLoc),
		Direction:     byte(msg.Direction),
		Hue:           uint16(msg.Hue),
		Flags:         byte(msg.Flags),
		NotorietyFlag: byte(msg.NotorietyFlags),
		Items:         make([]commands.UpdateObjectItem, 4),
	}
	// TODO remove hardcoded values
	command.Items[0] = commands.UpdateObjectItem{
		Serial:  0x40001FEE,
		Artwork: 0x0E75,
		Layer:   0x15,
		Hue:     0x0000,
	}
	command.Items[1] = commands.UpdateObjectItem{
		Serial:  0x40001FEC,
		Artwork: 0x203B,
		Layer:   0x0B,
		Hue:     0x044E,
	}
	command.Items[2] = commands.UpdateObjectItem{
		Serial:  0x40001FDA,
		Artwork: 0x204F,
		Layer:   0x16,
		Hue:     0x0022,
	}
	command.Items[3] = commands.UpdateObjectItem{
		Serial:  0x40001FD6,
		Artwork: 0x3EBE,
		Layer:   0x19,
		Hue:     0x0000,
	}
	return command
}
