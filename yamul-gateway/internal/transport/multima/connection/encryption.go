package connection

import (
	"crypto/md5"
	"encoding/binary"
	"golang.org/x/crypto/twofish"
	"yamul-gateway/internal/dtos"
)

const (
	noEncryption = iota
	loginEncryption
	gameplayEncryption
)

const TwofishTableSize = 0x100
const Md5ResetFlag = 0xf

func detectEncryptionAlgorithm(buffer *InputDataBuffer, config *dtos.EncryptionConfig) error {
	if config.GameplayServer {
		config.EncryptionAlgorithm = gameplayEncryption
		return initializeGameplayEncryption(config)
	} else if packageIsLoginInClean(buffer) {
		config.EncryptionAlgorithm = noEncryption
	} else {
		config.EncryptionAlgorithm = loginEncryption
		initializeLoginEncryption(config)
	}
	return nil
}

func initializeGameplayEncryption(config *dtos.EncryptionConfig) error {
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
		CreateAnonymousLogger("encryption").Error(err)
		return err
	}

	config.TwofishCipher = cipher
	config.TwofishReset = 0
	config.TwofishTable = make([]byte, TwofishTableSize)
	for i := 0; i < TwofishTableSize; i++ {
		config.TwofishTable[i] = byte(i)
	}

	cache := make([]byte, 16)
	for i := 0; i < TwofishTableSize; i += 16 {
		cipher.Encrypt(cache, config.TwofishTable[i:])
		copy(config.TwofishTable[i:], cache)
	}

	// initialize md5 table
	config.Md5Position = 0
	config.Md5Digest = md5.Sum(config.TwofishTable)

	return nil
}

func initializeLoginEncryption(config *dtos.EncryptionConfig) {
	seed := config.Seed
	config.LoginKey0 = (((^seed) ^ 0x00001357) << 16) | ((seed ^ 0xFFFFAAAA) & 0x0000FFFF)
	config.LoginKey1 = ((seed ^ 0x43210000) >> 16) | (((^seed) ^ 0xABCDFFFF) & 0xFFFF0000)
	config.LoginClientKey0 = 0x3AE221ED
	config.LoginClientKey1 = 0xA9F47E7F
}

func packageIsLoginInClean(buffer *InputDataBuffer) bool {
	if buffer.incomingTcpData[buffer.offset] != 0x80 {
		return false
	}

	for i := 20; i < 31; i++ {
		if buffer.incomingTcpData[buffer.offset+i] != 0 {
			return false
		}
	}

	for i := 40; i < 51; i++ {
		if buffer.incomingTcpData[buffer.offset+i] != 0 {
			return false
		}
	}

	return true
}

func inputDecryption(buffer *InputDataBuffer, config *dtos.EncryptionConfig) {
	algorithm := getDecryptionAlgorithm(config)
	for i := buffer.offset; i < buffer.length; {
		i += algorithm(buffer.incomingTcpData[i:buffer.length], buffer.decryptedData[i:buffer.length])
	}
}

func getDecryptionAlgorithm(config *dtos.EncryptionConfig) func([]byte, []byte) int {
	if config.EncryptionAlgorithm == noEncryption {
		return noEncryptionAlgorithm
	}
	if config.EncryptionAlgorithm == loginEncryption {
		return decorateCryptologicFunction(config, loginDecryptionAlgorithm)
	}
	return decorateCryptologicFunction(config, gameplayDecryptionAlgorithm)
}

func gameplayDecryptionAlgorithm(config *dtos.EncryptionConfig, in byte) byte {
	if config.TwofishReset == TwofishTableSize {
		// reset table
		config.TwofishReset = 0
		cache := make([]byte, 16)
		for i := 0; i < TwofishTableSize; i += 16 {
			config.TwofishCipher.Encrypt(cache, config.TwofishTable[i:])
			copy(config.TwofishTable[i:], cache)
		}
	}
	out := in ^ config.TwofishTable[config.TwofishReset]
	config.TwofishReset++
	return out
}

func outputDecryption(buffer *OutputDataBuffer, config *dtos.EncryptionConfig) []byte {
	if config.EncryptionAlgorithm != gameplayEncryption {
		return buffer.decryptedData[:buffer.length]
	}

	compressedBytes := HuffManCompress(buffer.compressedData, buffer.decryptedData, buffer.length)
	for i := 0; i < compressedBytes; i++ {
		in := buffer.compressedData[i]
		out := in ^ config.Md5Digest[config.Md5Position]
		config.Md5Position = (config.Md5Position + 1) & Md5ResetFlag
		buffer.outgoingTcpData[i] = out
	}

	return buffer.outgoingTcpData[:compressedBytes]
}

func decorateCryptologicFunction(config *dtos.EncryptionConfig, algorithm func(config *dtos.EncryptionConfig, in byte) byte) func(in []byte, out []byte) int {
	return func(in []byte, out []byte) int {
		for i := 0; i < len(in); i++ {
			out[i] = algorithm(config, in[i])
		}
		return len(in)
	}
}

func loginDecryptionAlgorithm(config *dtos.EncryptionConfig, in byte) byte {
	// Decrypt the byte:
	result := byte(config.LoginKey0) ^ in

	oldK0, oldK1 := config.LoginKey0, config.LoginKey1

	config.LoginKey0 = ((oldK0 >> 1) | (oldK1 << 31)) ^ config.LoginClientKey1
	config.LoginKey1 = (((((oldK1 >> 1) | (oldK0 << 31)) ^ (config.LoginClientKey0 - 1)) >> 1) | (oldK0 << 31)) ^ config.LoginClientKey0
	return result
}

func noEncryptionAlgorithm(in []byte, out []byte) int {
	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}
	return len(in)
}
