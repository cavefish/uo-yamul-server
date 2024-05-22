package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func GeneralLightLevel(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.GeneralLightLevel(connection, mapToCommandGeneralLightLevel(msg.Body.GetGeneralLightLevel()))
}

func mapToCommandGeneralLightLevel(level *services.MsgGeneralLightLevel) commands.GeneralLightLevel {
	return commands.GeneralLightLevel{
		Level: byte(level.Level),
	}
}
