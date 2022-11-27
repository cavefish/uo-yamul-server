package handlers

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
	"yamul-gateway/utils/numbers"
)

func ListGameServers(client *connection.ClientConnection, response commands.ListGameServers) { // 0xa8
	client.StartPacket()
	defer client.EndPacket()

	packageLength := 6 + 40*len(response.Servers)

	client.WriteByte(0xA8)
	client.WriteUShort(uint16(packageLength))
	client.WriteByte(response.Flags)
	client.WriteUShort(uint16(len(response.Servers)))
	for i := 0; i < len(response.Servers); i++ {
		server := response.Servers[i]
		client.WriteUShort(uint16(i + 1))
		client.WriteFixedString(32, server.Name)
		client.WriteByte(server.PercentageOfPlayers)
		client.WriteByte(server.Timezone)
		ip, _ := addressToUInt(server.AddressIP)
		client.WriteUInt(numbers.BigEndianUInt32(ip))
	}

	logging.Debug("Sending server list %+v\n", response)
}
