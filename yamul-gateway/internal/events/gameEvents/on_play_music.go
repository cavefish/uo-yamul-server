package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func OnPlayMusic(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	command := toCommandPlayMusic(msg.Body.GetPlayMusic())
	handlers.PlayMusic(connection, command)
}

func toCommandPlayMusic(music *services.MsgPlayMusic) commands.PlayMusic {
	return commands.PlayMusic{
		MusicId: uint16(music.MusicId),
	}
}
