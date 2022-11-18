package commands

type CharactersStartLocation struct {
	Characters     []CharacterLogin
	StartingCities []StartingCity
	Flags          ClientFeatures
}

type CharacterLogin struct {
	Name     string
	Password string
}

type StartingCity struct {
	Name   string
	Tavern string
}
