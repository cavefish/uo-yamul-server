package multima

import (
	"net"
	"time"
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn, isGameplayServer bool) {
	client := connection.CreateConnectionHandler(conn, isGameplayServer)
	defer client.CloseConnection()

	if isGameplayServer {
		// In the second connection of the client, the encryption key is sent to the server first
		client.UpdateEncryptionSeed(client.ReadUInt())
	}

	go clientOutputBufferWorker(&client)
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

func clientOutputBufferWorker(client *connection.ClientConnection) {
	for !client.ShouldCloseConnection && client.Err == nil {
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
