package listeners

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

type CommandEvent[T any] struct {
	Client  *connection.ClientConnection
	Command T
}

type CommandListener[T any] func(event CommandEvent[T])

func Build[T any](clientConnection *connection.ClientConnection, command T) CommandEvent[T] {
	return CommandEvent[T]{
		clientConnection,
		command,
	}
}

func Trigger[T any](listener CommandListener[T], event CommandEvent[T]) {
	if listener == nil {
		onMissingListener(event)
		return
	}

	go listener(event)
}

func onMissingListener[T any](event CommandEvent[T]) {
	logging.Error("Missing listener %T", event.Command)
	event.Client.ShouldCloseConnection = true
}
