package mocks

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"yamul-gateway/internal/dtos"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/connection"
)

type ClientConnectionReadBufferMock struct {
	assert           *assert.Assertions
	buffer           []byte
	usedBufferLength int
	logger           interfaces.Logger
}

func (c *ClientConnectionReadBufferMock) Close() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) SendAnyData() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) ReceiveData() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) ProcessInputBuffer() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) ReadByte() byte {
	c.assert.Less(c.usedBufferLength, len(c.buffer))
	var b = c.buffer[c.usedBufferLength]
	c.usedBufferLength++
	return b
}

func (c *ClientConnectionReadBufferMock) WriteByte(value byte) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) ReadUShort() uint16 {
	var hi = c.ReadByte()
	var lo = c.ReadByte()
	return (uint16(hi) << 8) | uint16(lo)
}

func (c *ClientConnectionReadBufferMock) WriteUShort(value uint16) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) ReadUInt() uint32 {
	var b0 = c.ReadByte()
	var b1 = c.ReadByte()
	var b2 = c.ReadByte()
	var b3 = c.ReadByte()
	return uint32(b0)<<24 | uint32(b1)<<16 | uint32(b2)<<8 | uint32(b3)
}

func (c *ClientConnectionReadBufferMock) WriteUInt(value uint32) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) ReadFixedString(length int) string {
	return string(c.ReadFixedBytes(length))
}

func (c *ClientConnectionReadBufferMock) ReadFixedBytes(length int) []byte {
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = c.ReadByte()
	}
	return result
}

func (c *ClientConnectionReadBufferMock) WriteFixedString(length int, value string) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) UpdateEncryptionSeed(newSeed uint32) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) StartPacket() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) EndPacket() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) CheckEncryptionHandshake() {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) SetLogin(username string, password string) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) CreateGameConnection() error {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) KillConnection(err error) {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) IsConnectionHealthy() bool {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) GetLogger() interfaces.Logger {
	return c.logger
}

func (c *ClientConnectionReadBufferMock) GetStatus() *dtos.ClientConnectionStatus {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) GetEncryptionState() *dtos.EncryptionConfig {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) GetLoginDetails() *dtos.LoginDetails {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) GetConnection() net.Conn {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) GetGameService() interfaces.GameService {
	panic("Unimplemented mock behaviour")
}

func (c *ClientConnectionReadBufferMock) AssertBufferConsumed() {
	c.assert.EqualValues(len(c.buffer), c.usedBufferLength, "Buffer is fully consumed")
}

func CreateClientConnectionReadBufferMock(t *testing.T, buffer []byte) *ClientConnectionReadBufferMock {
	result := &ClientConnectionReadBufferMock{
		assert:           assert.New(t),
		buffer:           buffer,
		usedBufferLength: 0,
		logger:           connection.CreateAnonymousLogger(t.Name()),
	}
	var _ interfaces.ClientConnection = result
	return result
}
