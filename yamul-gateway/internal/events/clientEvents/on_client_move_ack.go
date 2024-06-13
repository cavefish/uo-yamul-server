package clientEvents

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnClientMoveAck(event listeners.CommandEvent[commands.MoveAck]) { // TODO
	event.Client.GetLogger().Infof("Received move ack from client: %v", event.Command)
}
