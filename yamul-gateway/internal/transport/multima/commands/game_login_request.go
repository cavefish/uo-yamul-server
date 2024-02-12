package commands

type GameLoginRequest struct {
	Username      string
	Password      string
	EncryptionKey uint32
}
