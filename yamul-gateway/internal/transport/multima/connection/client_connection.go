package connection

import (
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"net"
	"sync"
	"yamul-gateway/internal/dtos"
	"yamul-gateway/internal/interfaces"
	internalServices "yamul-gateway/internal/services/game"
)

func CreateConnectionHandler(conn net.Conn) interfaces.ClientConnection {
	encryptionConfig := dtos.EncryptionConfig{
		GameplayServer:      false,
		Seed:                0,
		EncryptionAlgorithm: noEncryption,
	}
	connection := &clientConnection{
		connection:            conn,
		shouldCloseConnection: false,
		inputBuffer:           CreateInputDataBuffer(),
		outputBuffer:          CreateOutputDataBuffer(),
		encryptionState:       encryptionConfig,
		status:                dtos.ClientConnectionStatus{},
		uuid:                  uuid.New().String(),
	}
	connection.logger = CreateConnectionLogger(fmt.Sprintf("clientConnection-%s", connection.uuid), connection)
	connection.logger.Debug("Connection started")
	return connection
}

type clientConnection struct {
	mutex                 sync.Mutex
	connection            net.Conn
	shouldCloseConnection bool
	inputBuffer           InputDataBuffer
	outputBuffer          OutputDataBuffer
	err                   error
	encryptionState       dtos.EncryptionConfig
	logger                interfaces.Logger
	status                dtos.ClientConnectionStatus
	loginDetails          dtos.LoginDetails
	gameService           interfaces.GameService
	uuid                  string
}

func (client *clientConnection) GetGameService() interfaces.GameService {
	return client.gameService
}

func (client *clientConnection) GetStatus() *dtos.ClientConnectionStatus {
	return &client.status
}

func (client *clientConnection) GetEncryptionState() *dtos.EncryptionConfig {
	return &client.encryptionState
}

func (client *clientConnection) GetLoginDetails() *dtos.LoginDetails {
	return &client.loginDetails
}

func (client *clientConnection) GetConnection() net.Conn {
	return client.connection
}

func (client *clientConnection) GetLogger() interfaces.Logger {
	return client.logger
}

func (client *clientConnection) decrypt() {
	inputDecryption(&client.inputBuffer, &client.encryptionState)
	client.logger.Debug(client.inputBuffer.printBuffer())
}

func (client *clientConnection) getOutputSlice() []byte {
	slice := outputDecryption(&client.outputBuffer, &client.encryptionState)
	client.logger.Debug(client.outputBuffer.printBuffer())
	return slice
}

func (client *clientConnection) Close() {
	_ = client.connection.Close()
}

func (client *clientConnection) sendDataIfAlmostFull(requiredSize int) error {
	buffer := &client.outputBuffer
	if buffer.length < BufferSize-requiredSize {
		return nil
	}
	return client.sendEverything()
}

func (client *clientConnection) SendAnyData() error {
	return client.sendEverything()
}

func (client *clientConnection) sendEverything() error {
	buffer := &client.outputBuffer
	if buffer.length == 0 {
		return nil
	}
	slice := client.getOutputSlice()
	sentLength, err := client.connection.Write(slice)
	if err != nil || sentLength != len(slice) {
		client.err = err
		return err
	}
	client.logger.Debugf("Sent %d encrypted bytes: %x", sentLength, slice)
	buffer.length = 0
	return nil
}

func (client *clientConnection) ReceiveData() error {
	if client.inputBuffer.offset < client.inputBuffer.length {
		return nil
	}
	// Read the incoming connection into the buffer.
	reqLen, err := client.connection.Read(client.inputBuffer.incomingTcpData)
	if client.shouldCloseConnection {
		client.inputBuffer.length = 0
		return nil
	}
	if err != nil {
		client.logger.Errorf("Error reading: %v", err.Error())
		client.err = err
		return err
	}
	client.inputBuffer.length = reqLen
	client.inputBuffer.offset = 0
	client.decrypt()
	return nil
}

func (client *clientConnection) ProcessInputBuffer() {
	commandCode := client.ReadByte()
	handler := ClientCommandHandlers[commandCode]
	handler(client, commandCode)
}

func (client *clientConnection) ReadByte() byte {
	if client.ReceiveData() != nil {
		return 0
	}
	value := client.inputBuffer.decryptedData[client.inputBuffer.offset]
	client.inputBuffer.offset++
	return value
}

func (client *clientConnection) WriteByte(value byte) {
	if client.sendDataIfAlmostFull(1) != nil {
		return
	}
	client.outputBuffer.decryptedData[client.outputBuffer.length] = value
	client.outputBuffer.length++
}

func (client *clientConnection) ReadUShort() uint16 {
	value := uint16(client.ReadByte())
	value = value<<8 | uint16(client.ReadByte())
	return value
}

func (client *clientConnection) WriteUShort(value uint16) {
	if client.sendDataIfAlmostFull(2) != nil {
		return
	}
	binary.BigEndian.PutUint16(client.outputBuffer.decryptedData[client.outputBuffer.length:], value)
	client.outputBuffer.length += 2
}

func (client *clientConnection) ReadUInt() uint32 {
	value := uint32(client.ReadByte())
	value = value<<8 | uint32(client.ReadByte())
	value = value<<8 | uint32(client.ReadByte())
	value = value<<8 | uint32(client.ReadByte())
	return value
}

func (client *clientConnection) WriteUInt(value uint32) {
	if client.sendDataIfAlmostFull(4) != nil {
		return
	}
	binary.BigEndian.PutUint32(client.outputBuffer.decryptedData[client.outputBuffer.length:], value)
	client.outputBuffer.length += 4
}

func (client *clientConnection) ReadFixedString(length int) string {
	values := string(client.ReadFixedBytes(length))
	i := len(values) - 1
	for ; i > 0; i-- {
		if values[i] != 0 {
			break
		}
	}
	return values[:i+1]
}

func (client *clientConnection) ReadFixedBytes(length int) []byte {
	value := make([]byte, length)
	for i := 0; i < length; i++ {
		value[i] = client.ReadByte()
	}
	return value
}

func (client *clientConnection) WriteFixedString(length int, value string) {
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

func (client *clientConnection) UpdateEncryptionSeed(newSeed uint32) {
	client.encryptionState.Seed = newSeed
	err := detectEncryptionAlgorithm(&client.inputBuffer, &client.encryptionState)
	if err == nil {
		client.decrypt()
	} else {
		client.err = err
	}
}

func (client *clientConnection) StartPacket() {
	client.mutex.Lock()
}

func (client *clientConnection) EndPacket() {
	_ = client.sendEverything()
	client.mutex.Unlock()
}

func (client *clientConnection) CheckEncryptionHandshake() {
	_ = client.ReceiveData()
	firstByte := client.inputBuffer.incomingTcpData[0]
	if firstByte&0x80 != 0 {
		client.logger.Infof("Connecting to login server: %x", firstByte)
		// High byte is unencrypted or basic encryption
		return
	}
	client.logger.Info("Connecting to game server")
	client.encryptionState.GameplayServer = true
	client.UpdateEncryptionSeed(client.ReadUInt())
}

func (client *clientConnection) SetLogin(username string, password string) {
	client.loginDetails = dtos.LoginDetails{
		Username:      username,
		Password:      password,
		CharacterSlot: 0,
	}
}

func (client *clientConnection) CreateGameConnection() error {
	ser, err := internalServices.CreateGameService(client)
	if err != nil {
		client.KillConnection(err)
		return err
	}
	client.gameService = ser
	return nil
}

func (client *clientConnection) KillConnection(err error) {
	if client.shouldCloseConnection == true {
		return
	}
	client.logger.SetLogField("fatal-error", err)
	client.err = err
	client.shouldCloseConnection = true
	client.logger.Error("Fatal error. Closing client connection.")
	if client.gameService != nil {
		client.gameService.Close()
	}
}

func (client *clientConnection) IsConnectionHealthy() bool {
	return client.shouldCloseConnection != true
}
