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
	name     string
}

func createHandler[T any](name string) *ListenerHandler[T] {
	return &ListenerHandler[T]{
		listener: nil,
		name:     name,
	}
}

func (handler *ListenerHandler[T]) Trigger(client interfaces.ClientConnection, body T) {
	event := CommandEvent[T]{
		client,
		body,
	}

	if handler.listener == nil {
		onMissingListener(handler, event)
		return
	}

	handler.listener(event)
}

func (handler *ListenerHandler[T]) SetListener(listener func(event CommandEvent[T])) {
	handler.listener = listener
}

var Listeners = struct {
	OnClientDoubleClick *ListenerHandler[uint32]
	OnClientMoveRequest *ListenerHandler[commands2.ClientMoveRequest]
	OnClientViewRange   *ListenerHandler[byte]
	OnGameLoginRequest  *ListenerHandler[commands2.GameLoginRequest]
	OnLoginRequest      *ListenerHandler[commands2.LoginRequestCommand]
	OnMoveAck           *ListenerHandler[commands2.MoveAck]
	OnOpenChatWindow    *ListenerHandler[string]
	OnPreLogin          *ListenerHandler[commands2.PreLogin]
	OnShardSelected     *ListenerHandler[commands2.ShardSelected]
}{
	OnClientDoubleClick: createHandler[uint32]("OnClientDoubleClick"),
	OnClientMoveRequest: createHandler[commands2.ClientMoveRequest]("OnClientMoveRequest"),
	OnClientViewRange:   createHandler[byte]("OnClientViewRange"),
	OnGameLoginRequest:  createHandler[commands2.GameLoginRequest]("OnGameLoginRequest"),
	OnLoginRequest:      createHandler[commands2.LoginRequestCommand]("OnLoginRequest"),
	OnMoveAck:           createHandler[commands2.MoveAck]("OnMoveAck"),
	OnOpenChatWindow:    createHandler[string]("OnOpenChatWindow"),
	OnPreLogin:          createHandler[commands2.PreLogin]("OnPreLogin"),
	OnShardSelected:     createHandler[commands2.ShardSelected]("OnShardSelected"),
} // TODO refactor to reduce duplicated code

func onMissingListener[T any](handler *ListenerHandler[T], event CommandEvent[T]) {
	event.Client.GetLogger().Errorf("Missing listener %s[%T]: %v", handler.name, event.Command, event.Command)
}
