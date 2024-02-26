package autoconfig

import (
	log "github.com/sirupsen/logrus"
	"yamul-gateway/internal/events"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/handlers"
)

func Setup() error {
	log.SetFormatter(&log.JSONFormatter{})

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
