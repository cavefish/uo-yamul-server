package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/utils/booleans"
)

var StatWindowInfo_PackageSize = []uint16{43, 64, 66, 68, 86, 89, 119}

func StatWindowInfo(client interfaces.ClientConnection, command commands.StatWindowInfo) { // 0x11
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x11)
	client.WriteUShort(StatWindowInfo_PackageSize[command.FlagDisplay])
	client.WriteUInt(command.CharacterID)
	client.WriteFixedString(30, command.CharacterName)
	client.WriteUShort(command.HitPointsCurrent)
	client.WriteUShort(command.HitPointsMax)
	client.WriteByte(booleans.BoolToByte(command.FlagNameAllowed))
	client.WriteByte(command.FlagDisplay)
	if command.FlagDisplay == commands.FlagDisplay_minimal {
		return
	}

	client.WriteByte(command.Gender)
	client.WriteUShort(command.Strength)
	client.WriteUShort(command.Intelligence)
	client.WriteUShort(command.StaminaCurrent)
	client.WriteUShort(command.StaminaMax)
	client.WriteUShort(command.ManaCurrent)
	client.WriteUShort(command.ManaMax)
	client.WriteUInt(command.Gold)
	client.WriteUShort(command.ResistancePhysical)
	client.WriteUShort(command.WeightCurrent)
	if command.FlagDisplay == commands.FlagDisplay_normal {
		return
	}

	if command.FlagDisplay >= commands.FlagDisplay_stats_follwoers_resistances_weight {
		client.WriteUShort(command.WeightMax)
		client.WriteByte(command.Race)
	}

	client.WriteUShort(command.StatsCap)
	if command.FlagDisplay == commands.FlagDisplay_stats {
		return
	}

	client.WriteByte(command.FollowersCurrent)
	client.WriteByte(command.FollowersMax)
	if command.FlagDisplay == commands.FlagDisplay_stats_followers {
		return
	}

	client.WriteUShort(command.ResistanceFire)
	client.WriteUShort(command.ResistanceCold)
	client.WriteUShort(command.ResistancePoison)
	client.WriteUShort(command.ResistanceEnergy)
	client.WriteUShort(command.Luck)
	client.WriteUShort(command.DamageMin)
	client.WriteUShort(command.DamageMax)
	client.WriteUInt(command.TithingPoints)
	if command.FlagDisplay < commands.FlagDisplay_full_message {
		return
	}

	client.WriteUShort(command.ResistancePhysicalMax)
	client.WriteUShort(command.ResistanceFireMax)
	client.WriteUShort(command.ResistanceColdMax)
	client.WriteUShort(command.ResistancePoisonMax)
	client.WriteUShort(command.ResistanceEnergyMax)
	client.WriteUShort(command.DefenseChanceIncreaseCurrent)
	client.WriteUShort(command.DefenseChanceIncreaseMax)
	client.WriteUShort(command.HitChanceIncrease)
	client.WriteUShort(command.SwingSpeedIncrease)
	client.WriteUShort(command.DamageIncrease)
	client.WriteUShort(command.LowerReagentCost)
	client.WriteUShort(command.SpellDamageIncrease)
	client.WriteUShort(command.FasterCastRecovery)
	client.WriteUShort(command.FasterCasting)
	client.WriteUShort(command.LowerManaCost)

}
