package dtos

import "golang.org/x/crypto/twofish"

type EncryptionConfig struct {
	GameplayServer      bool
	Seed                uint32
	LoginKey0           uint32
	LoginKey1           uint32
	LoginClientKey0     uint32
	LoginClientKey1     uint32
	EncryptionAlgorithm int
	TwofishCipher       *twofish.Cipher
	TwofishReset        int
	TwofishTable        []byte
	Md5Digest           [16]byte
	Md5Position         int
	Version             string
}
