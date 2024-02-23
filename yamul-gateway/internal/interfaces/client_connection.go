package interfaces

import (
	"net"
	"yamul-gateway/internal/dtos"
)

type ClientConnection interface {
	Close()
	SendAnyData() error
	ReceiveData() error
	ProcessInputBuffer()
	ReadByte() byte
	WriteByte(value byte)
	ReadUShort() uint16
	WriteUShort(value uint16)
	ReadUInt() uint32
	WriteUInt(value uint32)
	ReadFixedString(length int) string
	ReadFixedBytes(length int) []byte
	WriteFixedString(length int, value string)
	UpdateEncryptionSeed(newSeed uint32)
	StartPacket()
	EndPacket()
	CheckEncryptionHandshake()
	SetLogin(username string, password string)
	CreateGameConnection() error
	KillConnection(err error)
	IsConnectionHealthy() bool
	GetLogger() Logger
	GetStatus() *dtos.ClientConnectionStatus
	GetEncryptionState() *dtos.EncryptionConfig
	GetLoginDetails() *dtos.LoginDetails
	GetConnection() net.Conn
	GetGameService() GameService
}
