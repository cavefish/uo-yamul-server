package connection

import (
	"net"
	"sync"
	"yamul-gateway/internal/logging"
)

const BufferSize = 1024

func CreateConnectionHandler(conn net.Conn) ClientConnection {
	inputBuffer := DataBuffer{
		rawData:       make([]byte, BufferSize),
		decryptedData: make([]byte, BufferSize),
	}
	outputBuffer := DataBuffer{
		rawData:       make([]byte, BufferSize),
		decryptedData: make([]byte, BufferSize),
	}
	return ClientConnection{
		Connection:   conn,
		inputBuffer:  inputBuffer,
		outputBuffer: outputBuffer,
	}
}

type ClientConnection struct {
	sync.Mutex
	Connection               net.Conn
	openingHandshakeReceived bool `default:"false"`
	ShouldCloseConnection    bool `default:"false"`
	inputBuffer              DataBuffer
	outputBuffer             DataBuffer
	Err                      error            `default:"nil"`
	EncryptSeed              EncryptionConfig `default:"nil"`
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

func (client *ClientConnection) CloseConnection() {
	_ = client.Connection.Close()
}

func (client *ClientConnection) sendDataIfAlmostFull(requiredSize int) error {
	if client.outputBuffer.length < BufferSize-requiredSize {
		return nil
	}
	err := client.encrypt()
	if err != nil {
		logging.Error("Error encrypting: ", err.Error())
		client.Err = err
		return err
	}
	buffer := client.outputBuffer
	sentLength, err := client.Connection.Write(buffer.rawData[buffer.offset:buffer.length])
	if err != nil || sentLength != buffer.length-buffer.offset {
		client.Err = err
		return err
	}
	buffer.offset = 0
	buffer.length = 0
	return nil
}

func (client *ClientConnection) SendAnyData() error {
	if client.outputBuffer.length == 0 {
		return nil
	}
	err := client.encrypt()
	if err != nil {
		logging.Error("Error encrypting: ", err.Error())
		client.Err = err
		return err
	}
	buffer := client.outputBuffer
	sentLength, err := client.Connection.Write(buffer.rawData[buffer.offset:buffer.length])
	if err != nil || sentLength != buffer.length-buffer.offset {
		client.Err = err
		return err
	}
	buffer.offset = 0
	buffer.length = 0
	return nil
}

func (client *ClientConnection) ReceiveData() error {
	if client.inputBuffer.offset < client.inputBuffer.length {
		return nil
	}
	// Read the incoming Connection into the buffer.
	reqLen, err := client.Connection.Read(client.inputBuffer.rawData)
	if client.ShouldCloseConnection {
		client.inputBuffer.length = 0
		return nil
	}
	if err != nil {
		logging.Error("Error reading: ", err.Error())
		client.Err = err
		return err
	}
	client.inputBuffer.length = reqLen
	client.inputBuffer.offset = 0
	err = client.decrypt()
	if err != nil {
		logging.Error("Error decrypting: ", err.Error())
		client.Err = err
		return err
	}
	client.inputBuffer.printBuffer()
	return nil
}

func (client *ClientConnection) ProcessInputBuffer() {
	commandCode := client.ReadByte()
	logging.Debug("Processing command %X\n", commandCode)
	handler := ClientCommandHandlers[commandCode]
	handler(client, commandCode)
}

func (client *ClientConnection) ReadByte() byte {
	if client.ReceiveData() != nil {
		return 0
	}
	value := client.inputBuffer.decryptedData[client.inputBuffer.offset]
	client.inputBuffer.offset++
	return value
}

func (client *ClientConnection) WriteByte(value byte) {
	if client.sendDataIfAlmostFull(1) != nil {
		return
	}
	client.outputBuffer.decryptedData[client.outputBuffer.length] = value
	client.outputBuffer.length++
}

func (client *ClientConnection) ReadInt() int32 {
	value := int32(client.ReadByte())
	value = value<<8 | int32(client.ReadByte())
	value = value<<8 | int32(client.ReadByte())
	value = value<<8 | int32(client.ReadByte())
	return value
}

func (client *ClientConnection) ReadFixedString(length int) string {
	value := make([]byte, length)
	for i := 0; i < length; i++ {
		value[i] = client.ReadByte()
	}
	return string(value)
}
