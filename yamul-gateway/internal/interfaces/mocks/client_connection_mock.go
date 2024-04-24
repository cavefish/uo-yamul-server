package mocks

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"yamul-gateway/internal/dtos"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/connection"
)

type ClientConnectionMock struct {
	assert           *assert.Assertions
	mutexIsLocked    bool
	mutexAlreadyUsed bool
	buffer           []byte
	usedBufferLength int
	logger           interfaces.Logger
}

func (c *ClientConnectionMock) Close() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) SendAnyData() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) ReceiveData() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) ProcessInputBuffer() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) ReadByte() byte {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) WriteByte(value byte) {
	c.assert.True(c.mutexIsLocked, "Mutex not locked")
	c.assert.Less(c.usedBufferLength, 1024, "Buffer overflow. Message too large.")
	c.buffer[c.usedBufferLength] = value
	c.usedBufferLength++
}

func (c *ClientConnectionMock) ReadUShort() uint16 {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) WriteUShort(value uint16) {
	c.WriteByte(byte(value >> 8))
	c.WriteByte(byte(value))
}

func (c *ClientConnectionMock) ReadUInt() uint32 {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) WriteUInt(value uint32) {
	c.WriteByte(byte(value >> 24))
	c.WriteByte(byte(value >> 16))
	c.WriteByte(byte(value >> 8))
	c.WriteByte(byte(value))
}

func (c *ClientConnectionMock) ReadFixedString(length int) string {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) ReadFixedBytes(length int) []byte {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) WriteFixedString(length int, value string) {
	for i := 0; i < length; i++ {
		if i < len(value) {
			c.WriteByte(value[i])
		} else {
			c.WriteByte(0)
		}
	}
}

func (c *ClientConnectionMock) UpdateEncryptionSeed(newSeed uint32) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) StartPacket() {
	c.assert.False(c.mutexIsLocked, "Mutex is already Locked")
	c.assert.False(c.mutexAlreadyUsed, "Mutex was already Locked")
	c.mutexIsLocked = true
}

func (c *ClientConnectionMock) EndPacket() {
	c.assert.True(c.mutexIsLocked, "Mutex is not Locked")
	c.mutexIsLocked = false
	c.mutexAlreadyUsed = true
}

func (c *ClientConnectionMock) CheckEncryptionHandshake() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) SetLogin(username string, password string) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) CreateGameConnection() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) KillConnection(err error) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) IsConnectionHealthy() bool {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) GetLogger() interfaces.Logger {
	return c.logger
}

func (c *ClientConnectionMock) GetStatus() *dtos.ClientConnectionStatus {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) GetEncryptionState() *dtos.EncryptionConfig {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) GetLoginDetails() *dtos.LoginDetails {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) GetConnection() net.Conn {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) GetGameService() interfaces.GameService {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionMock) AssertSentBuffer(expected []byte) {
	c.assert.False(c.mutexIsLocked, "Mutex is Locked")
	c.assert.EqualValues(expected, c.buffer[0:c.usedBufferLength])
}

func CreateClientConnectionMock(t *testing.T) *ClientConnectionMock {
	result := &ClientConnectionMock{
		assert:           assert.New(t),
		mutexIsLocked:    false,
		mutexAlreadyUsed: false,
		buffer:           make([]byte, 1024),
		usedBufferLength: 0,
		logger:           connection.CreateAnonymousLogger(t.Name()),
	}
	var _ interfaces.ClientConnection = result
	return result
}
