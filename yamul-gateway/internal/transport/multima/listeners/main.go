package listeners

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

var (
	OnClientDoubleClick    = createHandler[uint32]("OnClientDoubleClick")
	OnClientMoveAck        = createHandler[commands.MoveAck]("OnClientMoveAck")
	OnClientMoveRequest    = createHandler[commands.ClientMoveRequest]("OnClientMoveRequest")
	OnClientViewRange      = createHandler[byte]("OnClientViewRange")
	OnGameLoginRequest     = createHandler[commands.GameLoginRequest]("OnGameLoginRequest")
	OnLoginRequest         = createHandler[commands.LoginRequestCommand]("OnLoginRequest")
	OnOpenChatWindow       = createHandler[string]("OnOpenChatWindow")
	OnPreLogin             = createHandler[commands.PreLogin]("OnPreLogin")
	OnShardSelected        = createHandler[commands.ShardSelected]("OnShardSelected")
	OnUnicodeSpeechRequest = createHandler[commands.UnicodeSpeechSelected]("OnUnicodeSpeechRequest")
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

func onMissingListener[T any](handler *ListenerHandler[T], event CommandEvent[T]) {
	event.Client.GetLogger().Errorf("Missing listener %s[%T]: %v", handler.name, event.Command, event.Command)
}
