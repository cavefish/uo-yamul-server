package main

import (
	"net"
	"os"
	"sync"
	"yamul-gateway/internal/autoconfig"
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima"
)

const (
	SERVER_PORT = "2593"
	CONN_TYPE   = "tcp"
)

var listenerWg sync.WaitGroup

func main() {
	autoconfig.Setup()
	// Listen for incoming connections.
	listenerWg.Add(1)
	go listenToIncomingRequests(SERVER_PORT)
	listenerWg.Wait()
}

func listenToIncomingRequests(port string) {
	defer listenerWg.Done()
	l, err := net.Listen(CONN_TYPE, ":"+port)
	if err != nil {
		logging.Error("Error listening: %s\n", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	for {
		logging.Debug("Listening on port %s\n", port)
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			logging.Error("Error accepting connection %s\n", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go multima.ClientConnectionLoop(conn)
	}
}
