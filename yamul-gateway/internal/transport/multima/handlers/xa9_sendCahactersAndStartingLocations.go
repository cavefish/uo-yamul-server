package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func SendCharactersAndStartingLocations(client *connection.ClientConnection, body commands.CharactersStartLocation) { // 0xA9
	client.Lock()
	defer client.Unlock()

	size := 4 + 60*len(body.Characters) + 1 + 89*len(body.StartingCities) + 6

	client.WriteByte(0xA9)
	client.WriteUShort(uint16(size))

	client.WriteByte(byte(len(body.Characters)))
	for idx, character := range body.Characters {
		if idx > body.LastValidCharacter {
			client.WriteFixedString(60, "")
			continue
		}
		client.WriteFixedString(30, character.Name)
		client.WriteFixedString(30, character.Password)
	}

	client.WriteByte(byte(len(body.StartingCities)))
	for idx, city := range body.StartingCities {
		client.WriteByte(byte(idx))
		client.WriteFixedString(30, city.Name)
		client.WriteUShort(0)
		client.WriteFixedString(30, city.Tavern)
		client.WriteUShort(0)
		client.WriteUInt(city.CoordinateX)
		client.WriteUInt(city.CoordinateY)
		client.WriteUInt(city.CoordinateZ)
		client.WriteUInt(city.CoordinateMap)
		client.WriteUInt(city.ClilocDescription)
		client.WriteUInt(0)
	}

	client.WriteUInt(ConvertClientFeaturesToFlags(body.Flags))
	client.WriteUShort(uint16(body.LastValidCharacter))

}
