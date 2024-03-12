package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnHealthBarUpdate(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.HealthBarUpdate(connection, mapToHealthBarUpdate(msg.Body.GetHealthBar()))
}

func mapToHealthBarUpdate(body *services.MsgHealthBar) commands.HealthBarUpdate {
	values := make([]commands.HealthBarUpdateValues, len(body.Values))
	for idx := range body.Values {
		values[idx].Type = mapToHealthBarUpdateType(body.Values[idx].Type)
		values[idx].Enabled = body.Values[idx].Enabled
	}
	return commands.HealthBarUpdate{
		Serial: 0,
		Values: values,
	}
}

func mapToHealthBarUpdateType(valuesType services.MsgHealthBar_Values_Type) commands.HealthBarUpdateValuesType {
	if valuesType == services.MsgHealthBar_Values_GREEN {
		return commands.HealthBarUpdateValues_Green
	}
	return commands.HealthBarUpdateValues_Yellow
}
