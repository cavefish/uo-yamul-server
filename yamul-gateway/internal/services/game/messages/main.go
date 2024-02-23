package messages

import (
	gameService "yamul-gateway/backend/services"
	"yamul-gateway/internal/interfaces"
)

type MessageProcessor struct {
	Type gameService.MsgType
}

func (p MessageProcessor) Accept(connection interfaces.ClientConnection, msg *gameService.StreamPackage) {
	connection.GetLogger().Info("Processing message %s", msg.String())
}

var Processors = make(map[gameService.MsgType]MessageProcessor)
