package interfaces

import "yamul-gateway/backend/services"

type GameService interface {
	Close()
	Send(_type services.MsgType, message *services.Message)
}
