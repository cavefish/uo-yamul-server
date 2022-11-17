package listeners

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

type CommandEvent[T any] struct {
	Client  *connection.ClientConnection
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

func (handler *ListenerHandler[T]) Trigger(client *connection.ClientConnection, body T) {
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
	OnLoginRequest     *ListenerHandler[commands.LoginRequestCommand]
	OnShardSelected    *ListenerHandler[commands.ShardSelected]
	OnGameLoginRequest *ListenerHandler[commands.GameLoginRequest]
}{
	OnLoginRequest:     createHandler[commands.LoginRequestCommand](),
	OnShardSelected:    createHandler[commands.ShardSelected](),
	OnGameLoginRequest: createHandler[commands.GameLoginRequest](),
}

func onMissingListener[T any](event CommandEvent[T]) {
	logging.Error("Missing listener %T", event.Command)
}
