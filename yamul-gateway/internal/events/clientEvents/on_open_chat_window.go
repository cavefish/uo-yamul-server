package clientEvents

import (
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnOpenChatWindow(event listeners.CommandEvent[string]) { // TODO
	event.Client.GetLogger().Infof("Received open chat window from client: \"%s\"", event.Command)
}
