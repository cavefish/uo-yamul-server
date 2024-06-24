package commands

type HealthBarUpdate struct {
	Serial uint32
	Values []HealthBarUpdateValues
}

type HealthBarUpdateValuesType byte

const (
	HealthBarUpdateValues_Green  HealthBarUpdateValuesType = 1
	HealthBarUpdateValues_Yellow HealthBarUpdateValuesType = 2
)

type HealthBarUpdateValues struct {
	Type    HealthBarUpdateValuesType
	Enabled bool
}
