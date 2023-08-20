package multima

import (
	"net"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn) {
	client := connection.CreateConnectionHandler(conn)
	defer client.Close()

	client.CheckEncryptionHandshake()

	client.Logger.Info("Connection open")

	for !client.ShouldCloseConnection && client.Err == nil {
		client.ProcessInputBuffer()
		err := client.ReceiveData()
		if err != nil {
			client.Logger.Error("Error on connection loop %v", err)
			return
		}
	}

	if client.Err != nil {
		client.Logger.Error("error %v", client.Err)
	}
	client.Logger.Info("Connection closed")
}
