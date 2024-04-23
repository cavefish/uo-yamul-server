package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/utils/booleans"
)

func HealthBarUpdate(client interfaces.ClientConnection, command commands.HealthBarUpdate) { // 0x17
	client.StartPacket()
	defer client.EndPacket()

	var length uint16 = uint16(9 + len(command.Values)*3)

	client.WriteByte(0x17)
	client.WriteUShort(length)
	client.WriteUInt(command.Serial)
	client.WriteUShort(uint16(len(command.Values)))
	for idx := range command.Values {
		value := command.Values[idx]
		client.WriteUShort(uint16(value.Type))
		client.WriteByte(booleans.BoolToByte(value.Enabled) & 1)
	}
}
