package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func WarMode(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.Warmode(connection, msg.Body.GetWarmode().IsWarmode)
}
