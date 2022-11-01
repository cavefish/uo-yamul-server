package main

import (
	"fmt"
	"net"
	"sync"
)

const BufferSize = 1024

func createConnectionHandler(conn net.Conn) ClientConnection {
	inputBuffer := DataBuffer{
		rawData:       make([]byte, BufferSize),
		decryptedData: make([]byte, BufferSize),
	}
	outputBuffer := DataBuffer{
		rawData:       make([]byte, BufferSize),
		decryptedData: make([]byte, BufferSize),
	}
	return ClientConnection{
		connection:   conn,
		inputBuffer:  inputBuffer,
		outputBuffer: outputBuffer,
	}
}

type ClientConnection struct {
	connection               net.Conn
	openingHandshakeReceived bool `default:"false"`
	shouldCloseConnection    bool `default:"false"`
	inputBuffer              DataBuffer
	outputBuffer             DataBuffer
	outputMutex              sync.Mutex
	err                      error          `default:"nil"`
	encryptSeed              newSeedCommand `default:"nil"`
}

func (client *ClientConnection) decrypt() error {
	// Implementation without encryption
	if !client.openingHandshakeReceived {
		client.openingHandshakeReceived = true
		client.inputBuffer.decryptedData = client.inputBuffer.rawData
		client.outputBuffer.rawData = client.outputBuffer.decryptedData
	}
	return nil
}

func (client *ClientConnection) encrypt() error {
	// Implementation without encryption
	return nil
}

func (client *ClientConnection) closeConnection() {
	_ = client.connection.Close()
}

func (client *ClientConnection) sendDataIfAlmostFull(requiredSize int) error {
	if client.outputBuffer.length < BufferSize-requiredSize {
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
	buffer.offset = 0
	buffer.length = 0
	return nil
}

func (client *ClientConnection) sendAnyData() error {
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
	buffer.offset = 0
	buffer.length = 0
	return nil
}

func (client *ClientConnection) receiveData() error {
	if client.inputBuffer.offset < client.inputBuffer.length {
		return nil
	}
	// Read the incoming connection into the buffer.
	reqLen, err := client.connection.Read(client.inputBuffer.rawData)
	if client.shouldCloseConnection {
		client.inputBuffer.length = 0
		return nil
	}
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

func (client *ClientConnection) processInputBuffer() {
	commandCode := readByte(client)
	fmt.Printf("Processing command %X\n", commandCode)
	handler := clientCommandHandlers[commandCode]
	handler(client, commandCode)
}
