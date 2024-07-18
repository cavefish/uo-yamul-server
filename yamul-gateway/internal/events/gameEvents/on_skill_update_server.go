package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func SkillUpdateServer(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.SkillUpdateServer(connection, mapToSkillUpdateServerCommand(msg.Body.GetSkillUpdateServer()))
}

func mapToSkillUpdateServerCommand(input *services.MsgSkillUpdateServer) commands.SkillUpdateServer {
	output := commands.SkillUpdateServer{
		Type:   byte(input.Type),
		Skills: make([]*commands.SkillUpdateServerSkill, len(input.Skills)),
	}
	for idx, skill := range input.Skills {
		output.Skills[idx] = &commands.SkillUpdateServerSkill{
			Id:        uint16(skill.SkillId),
			Value:     uint16(skill.Value),
			BaseValue: uint16(skill.BaseValue),
			MaxValue:  uint16(skill.MaxValue),
			Status:    byte(skill.Status),
		}
	}
	return output
}
