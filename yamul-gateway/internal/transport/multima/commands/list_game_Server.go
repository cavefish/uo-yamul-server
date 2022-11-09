package commands

type ListGameServers struct {
	Flags   byte
	Servers []GameServer
}

type GameServer struct {
	Name                string
	PercentageOfPlayers byte
	Timezone            byte
	AddressIP           string
}
