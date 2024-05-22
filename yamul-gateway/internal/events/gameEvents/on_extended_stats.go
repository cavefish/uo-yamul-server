package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func ExtendedStats(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.ExtendedStats(connection, mapToCommandExtendedStats(msg.Body.GetExtendedStats()))
}

func mapToCommandExtendedStats(input *services.MsgExtendedStats) commands.ExtendedStats {
	switch msg := input.Msg.(type) {
	case *services.MsgExtendedStats_Dead:
		return commands.ExtendedStats{
			Type:     commands.ExtendedStats_Dead,
			ObjectID: msg.Dead.Id.Value,
			StrLock:  false,
			DexLock:  false,
			IntLock:  false,
			IsDead:   msg.Dead.IsDead,
		}
	case *services.MsgExtendedStats_Lock:
		return commands.ExtendedStats{
			Type:     commands.ExtendedStats_AttributeLock,
			ObjectID: msg.Lock.Id.Value,
			StrLock:  msg.Lock.StrLock,
			DexLock:  msg.Lock.DexLock,
			IntLock:  msg.Lock.IntLock,
			IsDead:   false,
		}
	}
	return commands.ExtendedStats{}
}
