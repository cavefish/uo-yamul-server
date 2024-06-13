package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func moveAckFromClient(client interfaces.ClientConnection) { // 0x22
	command := moveAckReadBuffer(client)

	listeners.OnClientMoveAck.Trigger(client, command)
}

func moveAckReadBuffer(client interfaces.ClientConnection) commands.MoveAck {
	sequence := client.ReadByte()
	status := client.ReadByte()
	return commands.MoveAck{
		Sequence: sequence,
		Status:   status,
	}
}

func MoveAckFromServer(client interfaces.ClientConnection, command commands.MoveAck) { // 0x22
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x22)
	client.WriteByte(command.Sequence)
	client.WriteByte(command.Status)
}
