package connection

type CommandHandler func(client *ClientConnection, commandCode byte)

var ClientCommandHandlers = make([]CommandHandler, 256)
