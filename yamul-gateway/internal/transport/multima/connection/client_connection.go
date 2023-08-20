package connection

import (
	"encoding/binary"
	"net"
	"sync"
)

func CreateConnectionHandler(conn net.Conn) *ClientConnection {
	encryptionConfig := EncryptionConfig{
		GameplayServer:      false,
		Seed:                0,
		encryptionAlgorithm: noEncryption,
	}
	connection := &ClientConnection{
		Connection:            conn,
		ShouldCloseConnection: false,
		inputBuffer:           CreateInputDataBuffer(),
		outputBuffer:          CreateOutputDataBuffer(),
		EncryptionState:       encryptionConfig,
		Status:                ClientConnectionStatus{},
	}
	connection.Logger = &logger{
		client:   connection,
		name:     "clientConnection",
		prefix:   "",
		logLevel: LogLevelDebug,
	}
	return connection
}

type ClientConnection struct {
	mutex                 sync.Mutex
	Connection            net.Conn
	ShouldCloseConnection bool
	inputBuffer           InputDataBuffer
	outputBuffer          OutputDataBuffer
	Err                   error
	EncryptionState       EncryptionConfig
	Logger                Logger
	Status                ClientConnectionStatus
}

type ClientConnectionStatus struct {
	UseMultiSight bool
}

func (client *ClientConnection) decrypt() {
	inputDecryption(&client.inputBuffer, &client.EncryptionState)
	client.Logger.Debug(client.inputBuffer.printBuffer())
}

func (client *ClientConnection) getOutputSlice() []byte {
	slice := outputDecryption(&client.outputBuffer, &client.EncryptionState)
	client.Logger.Debug(client.outputBuffer.printBuffer())
	return slice
}

func (client *ClientConnection) Close() {
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
	if buffer.length == 0 {
		return nil
	}
	slice := client.getOutputSlice()
	sentLength, err := client.Connection.Write(slice)
	if err != nil || sentLength != len(slice) {
		client.Err = err
		return err
	}
	client.Logger.Debug("Sent %d bytes", sentLength)
	buffer.length = 0
	return nil
}

func (client *ClientConnection) ReceiveData() error {
	if client.inputBuffer.offset < client.inputBuffer.length {
		return nil
	}
	// Read the incoming Connection into the buffer.
	reqLen, err := client.Connection.Read(client.inputBuffer.incomingTcpData)
	if client.ShouldCloseConnection {
		client.inputBuffer.length = 0
		return nil
	}
	if err != nil {
		client.Logger.Error("Error reading: %v", err.Error())
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
	return string(client.ReadFixedBytes(length))
}

func (client *ClientConnection) ReadFixedBytes(length int) []byte {
	value := make([]byte, length)
	for i := 0; i < length; i++ {
		value[i] = client.ReadByte()
	}
	return value
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

func (client *ClientConnection) CheckEncryptionHandshake() {
	_ = client.ReceiveData()
	firstByte := client.inputBuffer.incomingTcpData[0]
	if firstByte&0x80 != 0 {
		client.Logger.Info("Connecting to login server: %x", firstByte)
		// High byte is unencrypted or basic encryption
		return
	}
	client.Logger.Info("Connecting to game server")
	client.EncryptionState.GameplayServer = true
	client.UpdateEncryptionSeed(client.ReadUInt())
}
