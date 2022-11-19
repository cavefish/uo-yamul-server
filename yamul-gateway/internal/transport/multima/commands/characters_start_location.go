package commands

type CharactersStartLocation struct {
	Characters         []CharacterLogin
	LastValidCharacter int
	StartingCities     []StartingCity
	Flags              ClientFeatures
}

type CharacterLogin struct {
	Name     string
	Password string
}

type StartingCity struct {
	Name              string
	Tavern            string
	CoordinateX       uint32
	CoordinateY       uint32
	CoordinateZ       uint32
	CoordinateMap     uint32
	ClilocDescription uint32
}
