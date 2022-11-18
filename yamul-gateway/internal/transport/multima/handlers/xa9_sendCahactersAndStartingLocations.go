package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func SendCharactersAndStartingLocations(client *connection.ClientConnection, body commands.CharactersStartLocation) { // 0xA9
	client.Lock()
	defer client.Unlock()

	size := 9 + 60*len(body.Characters) + 61*len(body.StartingCities)

	client.WriteByte(0xA9)
	client.WriteUShort(uint16(size))

	client.WriteByte(byte(len(body.Characters)))
	for _, character := range body.Characters {
		client.WriteFixedString(30, character.Name)
		client.WriteFixedString(30, character.Password)
	}

	client.WriteByte(byte(len(body.StartingCities)))
	for idx, city := range body.StartingCities {
		client.WriteByte(byte(idx))
		client.WriteFixedString(30, city.Name)
		client.WriteFixedString(30, city.Tavern)
	}

	client.WriteUInt(convertClientFeaturesToFlags(body.Flags))

}
