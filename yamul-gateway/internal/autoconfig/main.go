package autoconfig

import (
	"yamul-gateway/internal/events"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/handlers"
)

func Setup() error {
	handlers.Setup()
	events.Setup()
	err := login.Setup()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	login.Close()
}
