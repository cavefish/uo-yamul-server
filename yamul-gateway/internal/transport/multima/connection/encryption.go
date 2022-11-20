package connection

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/twofish"
)

const (
	noEncryption = iota
	loginEncryption
	gameplayEncryption
)

const TwofishTableSize = 0x100
const Md5ResetFlag = 0xf

type EncryptionConfig struct {
	GameplayServer      bool
	Seed                uint32
	loginKey0           uint32
	loginKey1           uint32
	loginClientKey0     uint32
	loginClientKey1     uint32
	encryptionAlgorithm int
	twofishCipher       *twofish.Cipher
	twofishReset        int
	twofishTable        []byte
	md5Digest           [16]byte
	md5Position         int
	Version             string
}

func detectEncryptionAlgorithm(buffer *DataBuffer, config *EncryptionConfig) error {
	if config.GameplayServer {
		config.encryptionAlgorithm = gameplayEncryption
		return initializeGameplayEncryption(config)
	} else if packageIsLoginInClean(buffer) {
		config.encryptionAlgorithm = noEncryption
	} else {
		config.encryptionAlgorithm = loginEncryption
		initializeLoginEncryption(config)
	}
	return nil
}

func initializeGameplayEncryption(config *EncryptionConfig) error {
	seed := make([]uint32, 4)
	for i := 0; i < len(seed); i++ {
		seed[i] = config.Seed
	}

	seedAsBytes := make([]byte, len(seed)*4)
	for i := 0; i < len(seed); i++ {
		binary.BigEndian.PutUint32(seedAsBytes[i*4:], seed[i])
	}

	cipher, err := twofish.NewCipher(seedAsBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	config.twofishCipher = cipher
	config.twofishReset = 0
	config.twofishTable = make([]byte, TwofishTableSize)
	for i := 0; i < TwofishTableSize; i++ {
		config.twofishTable[i] = byte(i)
	}

	for i := 0; i < TwofishTableSize; i += 16 {
		cipher.Encrypt(config.twofishTable[i:], config.twofishTable[i:])
	}

	// initialize md5 table
	config.md5Position = 0
	config.md5Digest = md5.Sum(config.twofishTable)

	return nil
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
		return decorateCryptologicFunction(config, loginDecryptionAlgorithm)
	}
	return decorateCryptologicFunction(config, gameplayDecryptionAlgorithm)
}

func gameplayDecryptionAlgorithm(config *EncryptionConfig, in byte) byte {
	if config.twofishReset == TwofishTableSize {
		// reset table
		config.twofishReset = TwofishTableSize
	}
	out := in ^ config.twofishTable[config.twofishReset]
	config.twofishReset++
	return out
}

func outputDecryption(buffer *DataBuffer, config *EncryptionConfig) {
	algorithm := getEncryptionAlgorithm(config)
	for i := buffer.offset; i < buffer.length; {
		i += algorithm(buffer.decryptedData[i:buffer.length], buffer.rawData[i:buffer.length])
	}
}

func getEncryptionAlgorithm(config *EncryptionConfig) func(in []byte, out []byte) int {
	if config.encryptionAlgorithm == gameplayEncryption {
		return decorateCryptologicFunction(config, gameplayEncryptionAlgorithm)
	}
	return noEncryptionAlgorithm
}

func decorateCryptologicFunction(config *EncryptionConfig, algorithm func(config *EncryptionConfig, in byte) byte) func(in []byte, out []byte) int {
	return func(in []byte, out []byte) int {
		for i := 0; i < len(in); i++ {
			out[i] = algorithm(config, in[i])
		}
		return len(in)
	}
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

func gameplayEncryptionAlgorithm(config *EncryptionConfig, in byte) byte {
	out := in ^ config.md5Digest[config.md5Position]
	config.md5Position = (config.md5Position + 1) & Md5ResetFlag

	return out
}
