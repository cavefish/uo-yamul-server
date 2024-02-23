package listeners

import (
	commands2 "yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/connection"
)

type CommandEvent[T any] struct {
	Client  interfaces.ClientConnection
	Command T
}

type CommandListener[T any] func(event CommandEvent[T])

type ListenerHandler[T any] struct {
	listener CommandListener[T]
}

func createHandler[T any]() *ListenerHandler[T] {
	return &ListenerHandler[T]{
		listener: nil,
	}
}

func (handler *ListenerHandler[T]) Trigger(client interfaces.ClientConnection, body T) {
	event := CommandEvent[T]{
		client,
		body,
	}

	if handler.listener == nil {
		onMissingListener(event)
		return
	}

	handler.listener(event)
}

func (handler *ListenerHandler[T]) SetListener(listener func(event CommandEvent[T])) {
	handler.listener = listener
}

var Listeners = struct {
	OnLoginRequest     *ListenerHandler[commands2.LoginRequestCommand]
	OnShardSelected    *ListenerHandler[commands2.ShardSelected]
	OnGameLoginRequest *ListenerHandler[commands2.GameLoginRequest]
	OnPreLogin         *ListenerHandler[commands2.PreLogin]
}{
	OnLoginRequest:     createHandler[commands2.LoginRequestCommand](),
	OnShardSelected:    createHandler[commands2.ShardSelected](),
	OnGameLoginRequest: createHandler[commands2.GameLoginRequest](),
	OnPreLogin:         createHandler[commands2.PreLogin](),
}

func onMissingListener[T any](event CommandEvent[T]) {
	connection.LoggerFor("listeners").Error("Missing listener %T", event.Command)
}
