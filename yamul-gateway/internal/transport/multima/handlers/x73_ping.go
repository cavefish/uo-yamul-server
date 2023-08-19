package handlers

import "yamul-gateway/internal/transport/multima/connection"

func ping(client *connection.ClientConnection) { // 0x73
	ack := client.ReadByte()

	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x73)
	client.WriteByte(ack)
}
