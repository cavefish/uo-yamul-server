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

func (client *ClientConnection) decrypt() error {
	// Implementation without encryption
	if !client.openingHandshakeReceived {
		client.openingHandshakeReceived = true
		client.inputBuffer.offset = 21
		client.inputBuffer.decryptedData = client.inputBuffer.rawData
		client.outputBuffer.rawData = client.outputBuffer.decryptedData
	}
	return nil
}

func (client *ClientConnection) encrypt() error {
	// Implementation without encryption
	return nil
}

type ClientConnection struct {
	connection               net.Conn
	openingHandshakeReceived bool `default:"false"`
	shouldCloseConnection    bool `default:"false"`
	inputBuffer              DataBuffer
	outputBuffer             DataBuffer
	err                      error `default:"nil"`
}

func (client *ClientConnection) closeConnection() {
	_ = client.connection.Close()
}

func main() {
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
		err := client.receiveData()
		if err != nil {
			return
		}
		client.processInputBuffer()
		err = client.sendData()
		if err != nil {
			return
		}
	}
}

func (client *ClientConnection) sendData() error {
	if client.outputBuffer.length == 0 {
		return nil
	}
	err := client.encrypt()
	if err != nil {
		fmt.Println("Error encrypting: ", err.Error())
		client.err = err
		return err
	}
	buffer := client.outputBuffer
	sentLength, err := client.connection.Write(buffer.rawData[buffer.offset:buffer.length])
	if err != nil || sentLength != buffer.length-buffer.offset {
		client.err = err
		return err
	}
	return nil
}

func (client *ClientConnection) receiveData() error {
	if client.inputBuffer.offset < client.inputBuffer.length {
		return nil
	}
	// Read the incoming connection into the buffer.
	reqLen, err := client.connection.Read(client.inputBuffer.rawData)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		client.err = err
		return err
	}
	client.inputBuffer.length = reqLen
	client.inputBuffer.offset = 0
	err = client.decrypt()
	if err != nil {
		fmt.Println("Error decrypting: ", err.Error())
		client.err = err
		return err
	}
	printBuffer(client.inputBuffer)
	return nil
}

func createConnectionHandler(conn net.Conn) ClientConnection {
	inputBuffer := DataBuffer{
		rawData:       make([]byte, 1024),
		decryptedData: make([]byte, 1024),
	}
	outputBuffer := DataBuffer{
		rawData:       make([]byte, 1024),
		decryptedData: make([]byte, 1024),
	}
	return ClientConnection{
		connection:   conn,
		inputBuffer:  inputBuffer,
		outputBuffer: outputBuffer,
	}
}

func (client *ClientConnection) processInputBuffer() {
	commandCode := readByte(client)
	fmt.Printf("Processing command %X\n", commandCode)
}

func readByte(client *ClientConnection) byte {
	if client.receiveData() != nil {
		return 0
	}
	value := client.inputBuffer.decryptedData[client.inputBuffer.offset]
	client.inputBuffer.offset++
	return value
}

func printBuffer(buffer DataBuffer) {
	fmt.Printf("Buffer length %d\nraw:\t\t% x\n", buffer.length, buffer.rawData[0:buffer.length])
	fmt.Printf("decrypted:\t% x\n", buffer.decryptedData[0:buffer.length])
}
