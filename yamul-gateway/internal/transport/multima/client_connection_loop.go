package multima

import (
	"net"
	"time"
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn) {
	client := connection.CreateConnectionHandler(conn)
	defer client.CloseConnection()

	go clientOutputBufferWorker(&client)
	logging.Info("[%s]Connection open\n", conn.RemoteAddr())

	for !client.ShouldCloseConnection {
		client.ProcessInputBuffer()
		err := client.ReceiveData()
		if err != nil {
			logging.Error("[%s] Error on connection loop %v\n", conn.RemoteAddr(), err)
			return
		}
	}

	logging.Info("[%s] Connection closed\n", conn.RemoteAddr())
}

func clientOutputBufferWorker(client *connection.ClientConnection) {
	for !client.ShouldCloseConnection {
		time.Sleep(100 * time.Millisecond)
		client.Lock()
		err := client.SendAnyData()
		client.Unlock()
		if err != nil {
			logging.Error("[%s] Error on output worker loop %v\n", client.Connection.RemoteAddr(), err)
			return
		}
	}
}
