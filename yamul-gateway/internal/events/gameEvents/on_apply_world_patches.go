package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func ApplyWorldPatches(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := commands.WorldPatches{}
	handlers.WorldPatches(connection, command)
}
