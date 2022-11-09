package main

import (
	"fmt"
	"net"
	"os"
	"yamul-gateway/internal/autoconfig"
	"yamul-gateway/internal/transport/multima"
)

const (
	CONN_HOST = "192.168.0.60"
	CONN_PORT = "2593"
	CONN_TYPE = "tcp"
)

func main() {
	autoconfig.Setup()
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
		go multima.ClientConnectionLoop(conn)
	}
}
