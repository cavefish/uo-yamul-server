package multima

import (
	"fmt"
	"net"
	"time"
	"yamul-gateway/internal/transport/multima/connection"
)

// Handles incoming requests.
func ClientConnectionLoop(conn net.Conn) {
	client := connection.CreateConnectionHandler(conn)
	defer client.CloseConnection()

	go clientOutputBufferWorker(&client)
	fmt.Printf("Connection open %s\n", conn.RemoteAddr())

	for !client.ShouldCloseConnection {
		client.ProcessInputBuffer()
		err := client.ReceiveData()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Connection closed %s\n", conn.RemoteAddr())
}

func clientOutputBufferWorker(client *connection.ClientConnection) {
	for !client.ShouldCloseConnection {
		time.Sleep(100 * time.Millisecond)
		client.Lock()
		err := client.SendAnyData()
		client.Unlock()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
