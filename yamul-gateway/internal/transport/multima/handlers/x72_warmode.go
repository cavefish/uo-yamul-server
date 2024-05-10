package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func Warmode(client interfaces.ClientConnection, isWarmode bool) { // 0x72
	client.StartPacket()
	defer client.EndPacket()
	client.WriteByte(0x72)
	if isWarmode {
		client.WriteByte(1)
	} else {
		client.WriteByte(0)
	}
	client.WriteByte(0)
	client.WriteByte(0x32)
	client.WriteByte(0)
}
