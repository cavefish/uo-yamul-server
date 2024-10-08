package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"yamul-gateway/internal/autoconfig"
	"yamul-gateway/internal/transport/multima"
	"yamul-gateway/internal/transport/multima/connection"
)

const (
	SERVER_PORT = "2593"
	CONN_TYPE   = "tcp4"
)

var listenerWg sync.WaitGroup

func main() {
	err := autoconfig.Setup()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer autoconfig.Close()
	// Listen for incoming connections.
	listenerWg.Add(1)
	go listenToIncomingRequests(SERVER_PORT)
	listenerWg.Wait()
}

func listenToIncomingRequests(port string) {
	defer listenerWg.Done()

	logger := connection.CreateAnonymousLogger("main")

	l, err := net.Listen(CONN_TYPE, ":"+port)
	if err != nil {
		logger.Errorf("Error listening: %s", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	for {
		logger.Infof("Listening on port %s", port)
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			logger.Errorf("Error accepting connection %s", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go multima.ClientConnectionLoop(conn)
	}
}
