package connection

import (
	"yamul-gateway/internal/logging"
)

const BufferSize = 15360

func CreateInputDataBuffer() DataBuffer {
	return DataBuffer{
		rawData:        make([]byte, BufferSize),
		decryptedData:  make([]byte, BufferSize),
		compressedData: make([]byte, BufferSize),
		length:         0,
		offset:         0,
	}
}

func CreateOutputDataBuffer() DataBuffer {
	return DataBuffer{
		rawData:        make([]byte, BufferSize),
		decryptedData:  make([]byte, BufferSize),
		compressedData: make([]byte, BufferSize),
		length:         0,
		offset:         0,
	}
}

type DataBuffer struct {
	rawData        []byte
	decryptedData  []byte
	compressedData []byte
	length         int
	offset         int
}

func (buffer DataBuffer) printBuffer() {
	if buffer.offset >= buffer.length {
		return
	}
	logging.Debug("Buffer length %d\nraw:\t\t% x\n", buffer.length-buffer.offset, buffer.decryptedData[buffer.offset:buffer.length])
}
