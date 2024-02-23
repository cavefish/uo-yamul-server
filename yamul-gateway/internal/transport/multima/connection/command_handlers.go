package connection

import "yamul-gateway/internal/interfaces"

type CommandHandler func(client interfaces.ClientConnection, commandCode byte)

var ClientCommandHandlers = make([]CommandHandler, 256)
