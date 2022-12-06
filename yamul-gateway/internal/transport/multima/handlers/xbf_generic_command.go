package handlers

import "yamul-gateway/internal/transport/multima/connection"

func genericCommand(client *connection.ClientConnection) { // 0xBF
	size := client.ReadUShort() - 5
	subCommand := client.ReadUShort()
	body := client.ReadFixedBytes(int(size))

	client.Logger.Info("Received subcommand request %x", subCommand)
	client.Logger.Debug("size=%d [% x]", size, body)
}
