package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func clientMoveRequest(client interfaces.ClientConnection) { // 0x02
	direction := client.ReadByte()
	sequence := client.ReadByte()
	ackKey := client.ReadUInt()
	command := commands.ClientMoveRequest{
		Direction: direction,
		Sequence:  sequence,
		AckKey:    ackKey,
	}
	listeners.Listeners.OnClientMoveRequest.Trigger(client, command)
}
