package listeners

import (
	commands2 "yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
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
	OnGameLoginRequest *ListenerHandler[commands2.GameLoginRequest]
	OnLoginRequest     *ListenerHandler[commands2.LoginRequestCommand]
	OnMoveAck          *ListenerHandler[commands2.MoveAck]
	OnOpenChatWindow   *ListenerHandler[string]
	OnPreLogin         *ListenerHandler[commands2.PreLogin]
	OnShardSelected    *ListenerHandler[commands2.ShardSelected]
}{
	OnGameLoginRequest: createHandler[commands2.GameLoginRequest](),
	OnLoginRequest:     createHandler[commands2.LoginRequestCommand](),
	OnMoveAck:          createHandler[commands2.MoveAck](),
	OnOpenChatWindow:   createHandler[string](),
	OnPreLogin:         createHandler[commands2.PreLogin](),
	OnShardSelected:    createHandler[commands2.ShardSelected](),
}

func onMissingListener[T any](event CommandEvent[T]) {
	event.Client.GetLogger().Errorf("Missing listener %T: %v", event.Command, event.Command)
}
