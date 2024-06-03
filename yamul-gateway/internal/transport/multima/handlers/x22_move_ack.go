package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func moveAck(client interfaces.ClientConnection) { // 0x22
	command := moveAckReadBuffer(client)

	listeners.Listeners.OnMoveAck.Trigger(client, command)
}

func moveAckReadBuffer(client interfaces.ClientConnection) commands.MoveAck {
	sequence := client.ReadByte()
	status := client.ReadByte()
	return commands.MoveAck{
		Sequence: sequence,
		Status:   status,
	}
}
