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

	switch subCommand {
	case 0x0f:
		ignoredSubcommand(client, "Client info", body)
	default:
		client.GetLogger().Errorf("Unimplemented subcommand processor 0x%02x", subCommand)
	}
}

func ignoredSubcommand(client interfaces.ClientConnection, name string, body []byte) {
	client.GetLogger().Debugf("Received ignored sub command for %s: % x", name, body)
}
