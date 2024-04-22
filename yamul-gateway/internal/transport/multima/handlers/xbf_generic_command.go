package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func receiveGenericCommand(client interfaces.ClientConnection) { // 0xBF
	size := client.ReadUShort() - 5
	subCommand := client.ReadUShort()
	body := client.ReadFixedBytes(int(size))

	client.GetLogger().Infof("Received subcommand request 0x%02x", subCommand)
	client.GetLogger().Debugf("size=%d [% x]", size, body)

	client.GetLogger().Errorf("Unimplemented subcommand processor 0x%02x", subCommand)
}

func sendGenericCommand_changeWorldMap(client interfaces.ClientConnection, worldMapId byte) {
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0xBF)
	client.WriteUShort(5)
	client.WriteByte(worldMapId)
}
