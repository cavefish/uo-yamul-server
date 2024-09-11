package multima

import (
	"net"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn) {
	client := connection.CreateConnectionHandler(conn)
	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			client.GetLogger().Error(panicInfo)
		}
		client.Close()
	}()

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
