package connection

const (
	noEncryption = iota
	loginEncryption
	gameplayEncryption
)

type EncryptionConfig struct {
	GameplayServer      bool
	seed                uint32
	encryptionAlgorithm int
	VersionMajor        uint32
	VersionMinor        uint32
	VersionRevision     uint32
	VersionPatch        uint32
}

func detectEncryptionAlgorithm(buffer *DataBuffer, config *EncryptionConfig) {
	if config.GameplayServer {
		config.encryptionAlgorithm = gameplayEncryption
		panic("Unimplemented")
	} else if packageIsLoginInClean(buffer) {
		config.encryptionAlgorithm = noEncryption
	} else {
		config.encryptionAlgorithm = loginEncryption
	}
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
	for i := buffer.offset; i < buffer.length; i++ {
		buffer.decryptedData[i] = buffer.rawData[i]
	}
}

func outputDecryption(buffer *DataBuffer, config *EncryptionConfig) {
	for i := buffer.offset; i < buffer.length; i++ {
		buffer.rawData[i] = buffer.decryptedData[i]
	}
}
