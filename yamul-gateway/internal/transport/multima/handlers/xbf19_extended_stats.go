package handlers

import (
	"fmt"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/utils/booleans"
)

func ExtendedStats(client interfaces.ClientConnection, command commands.ExtendedStats) { // 0xbf19

	if command.Type == commands.ExtendedStats_Dead {
		extendedStatsDead(client, command)
	} else if command.Type == commands.ExtendedStats_AttributeLock {
		extendedStatsAttributeLock(client, command)
	} else {
		// TODO implement case 5
		client.KillConnection(fmt.Errorf("unknown command type %d for ExtendedStats", command.Type))
	}

}

func extendedStatsDead(client interfaces.ClientConnection, command commands.ExtendedStats) {
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0xbf)
	client.WriteUShort(11)
	client.WriteUShort(0x0019)
	client.WriteByte(commands.ExtendedStats_Dead)
	client.WriteUInt(command.ObjectID)
	client.WriteByte(booleans.BoolToByte(command.IsDead))
}

func extendedStatsAttributeLock(client interfaces.ClientConnection, command commands.ExtendedStats) {
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0xbf)
	client.WriteUShort(12)
	client.WriteUShort(0x0019)
	client.WriteByte(commands.ExtendedStats_AttributeLock)
	client.WriteUInt(command.ObjectID)
	client.WriteByte(0) // unused
	var lockFlags byte = 0
	if command.StrLock {
		lockFlags |= 0b00110000
	}
	if command.DexLock {
		lockFlags |= 0b00001100
	}
	if command.IntLock {
		lockFlags |= 0b00000011
	}
	client.WriteByte(lockFlags)
}
