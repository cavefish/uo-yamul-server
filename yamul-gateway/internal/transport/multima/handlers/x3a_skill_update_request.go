package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func SkillUpdateServer(client interfaces.ClientConnection, command commands.SkillUpdateServer) { // 0x3A
	client.StartPacket()
	defer client.EndPacket()

	hasCap := false
	switch command.Type {
	case 0x02:
	case 0x03:
	case 0xDF:
		hasCap = true
	}
	size := 4 + len(command.Skills)*7
	if hasCap {
		size = size + len(command.Skills)*2
	}

	client.WriteByte(0x3A)
	client.WriteUShort(uint16(size))
	client.WriteByte(command.Type)

	for _, skill := range command.Skills {
		client.WriteUShort(skill.Id)
		client.WriteUShort(skill.Value)
		client.WriteUShort(skill.BaseValue)
		client.WriteByte(skill.Status)
		if hasCap {
			client.WriteUShort(skill.MaxValue)
		}
	}

}
