package connection

import (
	"encoding/binary"
	"golang.org/x/crypto/twofish"
)

const (
	noEncryption = iota
	loginEncryption
	gameplayEncryption
)

type EncryptionConfig struct {
	GameplayServer      bool
	Seed                uint32
	loginKey0           uint32
	loginKey1           uint32
	loginClientKey0     uint32
	loginClientKey1     uint32
	encryptionAlgorithm int
	twofishCipher       *twofish.Cipher
	Version             string
}

func detectEncryptionAlgorithm(buffer *DataBuffer, config *EncryptionConfig) {
	if config.GameplayServer {
		config.encryptionAlgorithm = gameplayEncryption
		initializeGameplayEncryption(config)
	} else if packageIsLoginInClean(buffer) {
		config.encryptionAlgorithm = noEncryption
	} else {
		config.encryptionAlgorithm = loginEncryption
		initializeLoginEncryption(config)
	}
}

func initializeGameplayEncryption(config *EncryptionConfig) {
	seed := make([]byte, 4)
	binary.LittleEndian.PutUint32(seed, config.Seed)
	config.twofishCipher, _ = twofish.NewCipher(seed)
}

func initializeLoginEncryption(config *EncryptionConfig) {
	seed := config.Seed
	config.loginKey0 = (((^seed) ^ 0x00001357) << 16) | ((seed ^ 0xFFFFAAAA) & 0x0000FFFF)
	config.loginKey1 = ((seed ^ 0x43210000) >> 16) | (((^seed) ^ 0xABCDFFFF) & 0xFFFF0000)
	config.loginClientKey0 = 0x3AE221ED
	config.loginClientKey1 = 0xA9F47E7F
}

func packageIsLoginInClean(buffer *DataBuffer) bool {
	if buffer.rawData[buffer.offset] != 0x80 {
		return false
	}

	for i := 20; i < 31; i++ {
		if buffer.rawData[buffer.offset+i] != 0 {
			return false
		}
	}

	for i := 40; i < 51; i++ {
		if buffer.rawData[buffer.offset+i] != 0 {
			return false
		}
	}

	return true
}

func inputDecryption(buffer *DataBuffer, config *EncryptionConfig) {
	algorithm := getDecryptionAlgorithm(config)
	for i := buffer.offset; i < buffer.length; {
		i += algorithm(buffer.rawData[i:buffer.length], buffer.decryptedData[i:buffer.length])
	}
}

func getDecryptionAlgorithm(config *EncryptionConfig) func([]byte, []byte) int {
	if config.encryptionAlgorithm == noEncryption {
		return noEncryptionAlgorithm
	}
	if config.encryptionAlgorithm == loginEncryption {
		return func(in []byte, out []byte) int {
			for i := 0; i < len(in); i++ {
				out[i] = loginDecryptionAlgorithm(config, in[i])
			}
			return len(in)
		}
	}
	return func(in []byte, out []byte) int {
		return len(in)
	}
}

func outputDecryption(buffer *DataBuffer, config *EncryptionConfig) {
	algorithm := getEncryptionAlgorithm(config)
	for i := buffer.offset; i < buffer.length; {
		i += algorithm(buffer.decryptedData[i:buffer.length], buffer.rawData[i:buffer.length])
	}
}

func getEncryptionAlgorithm(config *EncryptionConfig) func(in []byte, out []byte) int {
	return noEncryptionAlgorithm
}

func loginDecryptionAlgorithm(config *EncryptionConfig, in byte) byte {
	// Decrypt the byte:
	result := byte(config.loginKey0) ^ in

	oldK0, oldK1 := config.loginKey0, config.loginKey1

	config.loginKey0 = ((oldK0 >> 1) | (oldK1 << 31)) ^ config.loginClientKey1
	config.loginKey1 = (((((oldK1 >> 1) | (oldK0 << 31)) ^ (config.loginClientKey0 - 1)) >> 1) | (oldK0 << 31)) ^ config.loginClientKey0
	return result
}

func noEncryptionAlgorithm(in []byte, out []byte) int {
	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}
	return len(in)
}
