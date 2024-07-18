package commands

type SkillUpdateServer struct {
	Type   byte
	Skills []*SkillUpdateServerSkill
}

type SkillUpdateServerSkill struct {
	Id        uint16
	Value     uint16
	BaseValue uint16
	MaxValue  uint16
	Status    byte
}
