package events

import (
	"yamul-gateway/internal/events/onLoginRequest"
	"yamul-gateway/internal/transport/multima/listeners"
)

func Setup() {
	listeners.Listeners.OnLoginRequest = onLoginRequest.OnLoginRequest
}
