package multima

import (
	"net"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn) {
	client := connection.CreateConnectionHandler(conn)
	defer client.CloseConnection()

	client.CheckEncryptionHandshake()

	client.Logger.Info("Connection open\n")

	for !client.ShouldCloseConnection && client.Err == nil {
		client.ProcessInputBuffer()
		err := client.ReceiveData()
		if err != nil {
			client.Logger.Error("Error on connection loop %v\n", err)
			return
		}
	}

	if client.Err != nil {
		client.Logger.Error("error %v\n", client.Err)
	}
	client.Logger.Info("Connection closed\n")
}
