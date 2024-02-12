package commands

type PreLogin struct {
	Name          string
	Password      string
	Slot          uint32
	EncryptionKey uint32
}
