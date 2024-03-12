package commands

type HealthBarUpdate struct {
	Serial uint32
	Values []HealthBarUpdateValues
}

type HealthBarUpdateValuesType byte

const (
	HealthBarUpdateValues_Green  HealthBarUpdateValuesType = 0
	HealthBarUpdateValues_Yellow HealthBarUpdateValuesType = 1
)

type HealthBarUpdateValues struct {
	Type    HealthBarUpdateValuesType
	Enabled bool
}
