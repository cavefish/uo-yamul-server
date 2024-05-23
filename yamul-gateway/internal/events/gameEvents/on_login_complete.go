package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnLoginComplete(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.LoginComplete(connection, nil)
}
