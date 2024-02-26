package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func receiveGenericCommand(client interfaces.ClientConnection) { // 0xBF
	size := client.ReadUShort() - 5
	subCommand := client.ReadUShort()
	body := client.ReadFixedBytes(int(size))

	client.GetLogger().Infof("Received subcommand request %x", subCommand)
	client.GetLogger().Debugf("size=%d [% x]", size, body)
}

func sendGenericCommand_changeWorldMap(client interfaces.ClientConnection, worldMapId byte) {
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xBF)
	client.WriteUShort(5)
	client.WriteByte(worldMapId)
}
