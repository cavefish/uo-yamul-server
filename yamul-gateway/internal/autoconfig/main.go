package autoconfig

import (
	"yamul-gateway/internal/events"
	"yamul-gateway/internal/transport/multima/handlers"
)

func Setup() {
	handlers.Setup()
	events.Setup()

}
