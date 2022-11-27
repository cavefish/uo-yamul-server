package multima

import (
	"net"
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn) {
	client := connection.CreateConnectionHandler(conn)
	defer client.CloseConnection()

	client.CheckEncryptionHandshake()

	logging.Info("[%s %s] Connection open\n", conn.LocalAddr(), conn.RemoteAddr())

	for !client.ShouldCloseConnection && client.Err == nil {
		client.ProcessInputBuffer()
		err := client.ReceiveData()
		if err != nil {
			logging.Error("[%s %s] Error on connection loop %v\n", conn.LocalAddr(), conn.RemoteAddr(), err)
			return
		}
	}

	if client.Err != nil {
		logging.Error("[%s %s] error %v\n", conn.LocalAddr(), conn.RemoteAddr(), client.Err)
	}
	logging.Info("[%s %s] Connection closed\n", conn.LocalAddr(), conn.RemoteAddr())
}
