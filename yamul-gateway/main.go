package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "192.168.0.60"
	CONN_PORT = "2593"
	CONN_TYPE = "tcp"
)

type DataBuffer struct {
	rawData       []byte
	decryptedData []byte
	length        int `default:"0"`
	offset        int `default:"0"`
}

func main() {
	setupCommandHandlers()
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go clientConnectionLoop(conn)
	}
}

// Handles incoming requests.
func clientConnectionLoop(conn net.Conn) {
	client := createConnectionHandler(conn)
	defer client.closeConnection()

	go clientOutputBufferWorker(&client)
	fmt.Printf("Connection open %s\n", conn.RemoteAddr())

	for !client.shouldCloseConnection {
		client.processInputBuffer()
		err := client.receiveData()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Connection closed %s\n", conn.RemoteAddr())
}

func clientOutputBufferWorker(client *ClientConnection) {
	for !client.shouldCloseConnection {
		time.Sleep(100 * time.Millisecond)
		client.outputMutex.Lock()
		err := client.sendAnyData()
		client.outputMutex.Unlock()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
