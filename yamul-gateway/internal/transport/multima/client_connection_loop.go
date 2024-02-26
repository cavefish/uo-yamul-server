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

	client.GetLogger().Infof("connection open")

	for client.IsConnectionHealthy() {
		client.ProcessInputBuffer()
		err := client.ReceiveData()
		if err != nil {
			client.GetLogger().Errorf("Error on connection loop %v", err)
			return
		}
	}

	client.GetLogger().Infof("connection closed")
}
