package commands

const (
	FlagDisplay_minimal                            = 0x00
	FlagDisplay_normal                             = 0x01
	FlagDisplay_stats                              = 0x02
	FlagDisplay_stats_followers                    = 0x03
	FlagDisplay_stats_followers_resistances        = 0x04
	FlagDisplay_stats_follwoers_resistances_weight = 0x05
	FlagDisplay_full_message                       = 0x06
)

type StatWindowInfo struct {
	CharacterID                  uint32
	CharacterName                string
	HitPointsCurrent             uint16
	HitPointsMax                 uint16
	FlagNameAllowed              bool
	FlagDisplay                  byte
	Gender                       byte
	Strength                     uint16
	Intelligence                 uint16
	StaminaCurrent               uint16
	StaminaMax                   uint16
	ManaCurrent                  uint16
	ManaMax                      uint16
	Gold                         uint32
	ResistancePhysical           uint16
	WeightCurrent                uint16
	WeightMax                    uint16
	Race                         byte
	StatsCap                     uint16
	FollowersCurrent             byte
	FollowersMax                 byte
	ResistanceFire               uint16
	ResistanceCold               uint16
	ResistancePoison             uint16
	ResistanceEnergy             uint16
	Luck                         uint16
	DamageMin                    uint16
	DamageMax                    uint16
	TithingPoints                uint32
	ResistancePhysicalMax        uint16
	ResistanceFireMax            uint16
	ResistanceColdMax            uint16
	ResistancePoisonMax          uint16
	ResistanceEnergyMax          uint16
	DefenseChanceIncreaseCurrent uint16
	DefenseChanceIncreaseMax     uint16
	HitChanceIncrease            uint16
	SwingSpeedIncrease           uint16
	DamageIncrease               uint16
	LowerReagentCost             uint16
	SpellDamageIncrease          uint16
	FasterCastRecovery           uint16
	FasterCasting                uint16
	LowerManaCost                uint16
}
