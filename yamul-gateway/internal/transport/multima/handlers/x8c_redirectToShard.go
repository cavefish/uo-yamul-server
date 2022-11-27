package handlers

import (
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

func RedirectToShard(client *connection.ClientConnection, body commands.RedirectToShard) { // 0x8C
	client.StartPacket()
	defer client.EndPacket()

	ip, port := addressToUInt(body.AddressIP)

	client.WriteByte(0x8C)
	client.WriteUInt(ip)
	client.WriteUShort(port)
	client.WriteUInt(body.EncryptionKey & 0x3fffffff)
}
