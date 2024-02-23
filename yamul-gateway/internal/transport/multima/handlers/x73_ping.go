package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func ping(client interfaces.ClientConnection) { // 0x73
	ack := client.ReadByte()

	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x73)
	client.WriteByte(ack)
}
