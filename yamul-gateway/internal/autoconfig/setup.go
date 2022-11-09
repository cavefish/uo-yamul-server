package autoconfig

import "yamul-gateway/internal/transport/multima/handlers"

func Setup() {
	handlers.SetupCommandHandlers()

}
