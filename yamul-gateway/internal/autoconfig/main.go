package autoconfig

import (
	log "github.com/sirupsen/logrus"
	"yamul-gateway/internal/events"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/handlers"
)

func Setup() error {
	formatter := &log.TextFormatter{}
	formatter.ForceColors = true
	log.SetFormatter(formatter)
	//log.SetLevel(log.DebugLevel)

	handlers.Setup()
	events.Setup()
	err := login.Module.Setup()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	login.Module.Close()
}
