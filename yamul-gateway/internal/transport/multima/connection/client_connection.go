package connection

import (
	"encoding/binary"
	"net"
	"sync"
	"yamul-gateway/internal/logging"
)

func CreateConnectionHandler(conn net.Conn, isGameplayServer bool) ClientConnection {
	encryptionConfig := EncryptionConfig{
		GameplayServer:      isGameplayServer,
		Seed:                0,
		encryptionAlgorithm: noEncryption,
	}
	return ClientConnection{
		Connection:            conn,
		ShouldCloseConnection: false,
		inputBuffer:           CreateInputDataBuffer(),
		outputBuffer:          CreateOutputDataBuffer(),
		EncryptionState:       encryptionConfig,
	}
}

type ClientConnection struct {
	mutex                 sync.Mutex
	Connection            net.Conn
	ShouldCloseConnection bool
	inputBuffer           DataBuffer
	outputBuffer          DataBuffer
	Err                   error
	EncryptionState       EncryptionConfig
}

func (client *ClientConnection) decrypt() {
	inputDecryption(&client.inputBuffer, &client.EncryptionState)
	client.inputBuffer.printBuffer()
}

func (client *ClientConnection) getOutputSlice() []byte {
	slice := outputDecryption(&client.outputBuffer, &client.EncryptionState)
	client.outputBuffer.printBuffer()
	return slice
}

func (client *ClientConnection) CloseConnection() {
	_ = client.Connection.Close()
}

func (client *ClientConnection) sendDataIfAlmostFull(requiredSize int) error {
	buffer := &client.outputBuffer
	if buffer.length < BufferSize-requiredSize {
		return nil
	}
	return client.sendEverything()
}

func (client *ClientConnection) SendAnyData() error {
	return client.sendEverything()
}

func (client *ClientConnection) sendEverything() error {
	buffer := &client.outputBuffer
	bytesToSend := buffer.length - buffer.offset
	if bytesToSend == 0 {
		return nil
	}
	slice := client.getOutputSlice()
	sentLength, err := client.Connection.Write(slice)
	if err != nil || sentLength != len(slice) {
		client.Err = err
		return err
	}
	logging.Debug("Sent %d bytes\n", sentLength)
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
		logging.Error("Error reading: %v\n", err.Error())
		client.Err = err
		return err
	}
	client.inputBuffer.length = reqLen
	client.inputBuffer.offset = 0
	client.decrypt()
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

func (client *ClientConnection) ReadUShort() uint16 {
	value := uint16(client.ReadByte())
	value = value<<8 | uint16(client.ReadByte())
	return value
}

func (client *ClientConnection) WriteUShort(value uint16) {
	if client.sendDataIfAlmostFull(2) != nil {
		return
	}
	binary.BigEndian.PutUint16(client.outputBuffer.decryptedData[client.outputBuffer.length:], value)
	client.outputBuffer.length += 2
}

func (client *ClientConnection) ReadUInt() uint32 {
	value := uint32(client.ReadByte())
	value = value<<8 | uint32(client.ReadByte())
	value = value<<8 | uint32(client.ReadByte())
	value = value<<8 | uint32(client.ReadByte())
	return value
}

func (client *ClientConnection) WriteUInt(value uint32) {
	if client.sendDataIfAlmostFull(4) != nil {
		return
	}
	binary.BigEndian.PutUint32(client.outputBuffer.decryptedData[client.outputBuffer.length:], value)
	client.outputBuffer.length += 4
}

func (client *ClientConnection) ReadFixedString(length int) string {
	value := make([]byte, length)
	for i := 0; i < length; i++ {
		value[i] = client.ReadByte()
	}
	return string(value)
}

func (client *ClientConnection) WriteFixedString(length int, value string) {
	if client.sendDataIfAlmostFull(length) != nil {
		return
	}
	limit := len(value)
	if limit >= length {
		limit = length - 1
	}
	for i := 0; i < limit; i++ {
		client.WriteByte(value[i])
	}
	for i := limit; i < length; i++ {
		client.WriteByte(0x00)
	}
}

func (client *ClientConnection) UpdateEncryptionSeed(newSeed uint32) {
	client.EncryptionState.Seed = newSeed
	err := detectEncryptionAlgorithm(&client.inputBuffer, &client.EncryptionState)
	if err == nil {
		client.decrypt()
	} else {
		client.Err = err
	}
}

func (client *ClientConnection) StartPacket() {
	client.mutex.Lock()
}

func (client *ClientConnection) EndPacket() {
	_ = client.sendEverything()
	client.mutex.Unlock()
}
