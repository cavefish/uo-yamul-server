package messages

import (
	gameService "yamul-gateway/backend/services"
	"yamul-gateway/internal/interfaces"
)

type Processor func(connection interfaces.ClientConnection, msg *gameService.StreamPackage)

type MessageProcessor struct {
	Type   gameService.MsgType
	method Processor
}

func (p MessageProcessor) Accept(connection interfaces.ClientConnection, msg *gameService.StreamPackage) {
	connection.GetLogger().Infof("Processing message %s", msg.String())
	p.method(connection, msg)
}

func UnimplementedProcessor(connection interfaces.ClientConnection, msg *gameService.StreamPackage) {
	connection.GetLogger().Warningf("Unimplemented processor %s", msg.String())
}

var Processors = make(map[gameService.MsgType]MessageProcessor)

func RegisterProcessor(msgType gameService.MsgType, processor Processor) {
	Processors[msgType] = MessageProcessor{
		Type:   msgType,
		method: processor,
	}
}
