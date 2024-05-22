package gameEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/handlers"
)

func StatWindowInfo(connection interfaces.ClientConnection, msg *services.StreamPackage) {
	handlers.StatWindowInfo(connection, mapToCommandStatWindowInfo(msg.Body.GetStatWindowInfo()))
}

func mapToCommandStatWindowInfo(level1 *services.MsgStatWindowInfo) commands.StatWindowInfo {
	c := commands.StatWindowInfo{}
	c.CharacterID = level1.CharacterID.Value
	c.CharacterName = level1.CharacterName
	c.HitPointsCurrent = uint16(level1.HitPointsCurrent)
	c.HitPointsMax = uint16(level1.HitPointsMax)
	c.FlagNameAllowed = level1.FlagNameAllowed
	c.FlagDisplay = byte(level1.FlagDisplay)
	level2 := level1.Level2
	if level2 == nil {
		return c
	}
	c.Gender = byte(level2.Gender)
	c.Strength = uint16(level2.Strength)
	c.Intelligence = uint16(level2.Intelligence)
	c.StaminaCurrent = uint16(level2.StaminaCurrent)
	c.StaminaMax = uint16(level2.StaminaMax)
	c.ManaCurrent = uint16(level2.ManaCurrent)
	c.ManaMax = uint16(level2.ManaMax)
	c.Gold = level2.Gold
	c.ResistancePhysical = uint16(level2.ResistancePhysical)
	c.WeightCurrent = uint16(level2.WeightCurrent)
	level3 := level2.Level3
	if level3 == nil {
		return c
	}
	c.StatsCap = uint16(level3.StatsCap)
	level4 := level3.Level4
	if level4 == nil {
		return c
	}
	c.FollowersCurrent = byte(level4.FollowersCurrent)
	c.FollowersMax = byte(level4.FollowersMax)
	level5 := level4.Level5
	if level5 == nil {
		return c
	}
	c.ResistanceFire = uint16(level5.ResistanceFire)
	c.ResistanceCold = uint16(level5.ResistanceCold)
	c.ResistancePoison = uint16(level5.ResistancePoison)
	c.ResistanceEnergy = uint16(level5.ResistanceEnergy)
	c.Luck = uint16(level5.Luck)
	c.DamageMin = uint16(level5.DamageMin)
	c.DamageMax = uint16(level5.DamageMax)
	c.TithingPoints = level5.TithingPoints
	level6 := level5.Level6
	if level6 == nil {
		return c
	}
	c.WeightMax = uint16(level6.WeightMax)
	c.Race = byte(level6.Race)
	level7 := level6.Level7
	if level7 == nil {
		return c
	}
	c.ResistancePhysicalMax = uint16(level7.ResistancePhysicalMax)
	c.ResistanceFireMax = uint16(level7.ResistanceFireMax)
	c.ResistanceColdMax = uint16(level7.ResistanceColdMax)
	c.ResistancePoisonMax = uint16(level7.ResistancePoisonMax)
	c.ResistanceEnergyMax = uint16(level7.ResistanceEnergyMax)
	c.DefenseChanceIncreaseCurrent = uint16(level7.DefenseChanceIncreaseCurrent)
	c.DefenseChanceIncreaseMax = uint16(level7.DefenseChanceIncreaseMax)
	c.HitChanceIncrease = uint16(level7.HitChanceIncrease)
	c.SwingSpeedIncrease = uint16(level7.SwingSpeedIncrease)
	c.DamageIncrease = uint16(level7.DamageIncrease)
	c.LowerReagentCost = uint16(level7.LowerReagentCost)
	c.SpellDamageIncrease = uint16(level7.SpellDamageIncrease)
	c.FasterCastRecovery = uint16(level7.FasterCastRecovery)
	c.FasterCasting = uint16(level7.FasterCasting)
	c.LowerManaCost = uint16(level7.LowerManaCost)
	return c
}
