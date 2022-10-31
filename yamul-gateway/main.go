package main

import (
	"fmt"
	"net"
	"os"
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

	for !client.shouldCloseConnection {
		fmt.Println("Waiting for data")
		err := client.receiveData()
		if err != nil {
			fmt.Println(err)
			return
		}
		client.processInputBuffer()
		err = client.sendAnyData()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
