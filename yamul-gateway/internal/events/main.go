package events

import (
	"yamul-gateway/internal/events/onLoginRequest"
	"yamul-gateway/internal/events/onShardSelected"
	"yamul-gateway/internal/transport/multima/listeners"
)

func Setup() {
	listeners.Listeners.OnLoginRequest.SetListener(onLoginRequest.OnLoginRequest)
	listeners.Listeners.OnShardSelected.SetListener(onShardSelected.OnShardSelected)
}
