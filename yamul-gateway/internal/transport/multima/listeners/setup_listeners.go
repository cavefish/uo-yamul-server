package listeners

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/commands/x82_loginDenied"
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
	response := x82_loginDenied.LoginDeniedCommand{
		Reason: x82_loginDenied.CommunicationProblem,
	}
	logging.Error("Missing listener %T", event.Command)
	x82_loginDenied.LoginDenied(event.Client, response)
}
