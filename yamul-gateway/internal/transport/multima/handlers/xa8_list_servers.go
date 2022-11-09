package handlers

import (
	"strconv"
	"strings"
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func ListGameServers(client *connection.ClientConnection, response commands.ListGameServers) { // 0xa8
	client.Lock()
	defer client.Unlock()

	packageLength := 6 + 40*len(response.Servers)

	client.WriteByte(0xA8)
	client.WriteUShort(uint16(packageLength))
	client.WriteByte(response.Flags)
	client.WriteUShort(uint16(len(response.Servers)))
	for i := 0; i < len(response.Servers); i++ {
		server := response.Servers[i]
		client.WriteUShort(uint16(i))
		client.WriteFixedString(32, server.Name)
		client.WriteByte(server.PercentageOfPlayers)
		client.WriteByte(server.Timezone)
		client.WriteUInt(addressToUInt(server.AddressIP))
	}

	logging.Debug("Sending server list %+v\n", response)
}

func addressToUInt(value string) uint32 {
	ip := strings.Split(value, ":")
	ipTokens := strings.Split(ip[0], ".")
	var result uint32 = 0
	for i := 0; i < 4; i++ {
		v, _ := strconv.Atoi(ipTokens[3-i])
		result = result<<8 | uint32(v)
	}
	return result
}
