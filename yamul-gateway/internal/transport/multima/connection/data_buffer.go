package connection

import (
	"yamul-gateway/internal/logging"
)

const BufferSize = 15360

func CreateDataBuffer() DataBuffer {
	return DataBuffer{
		rawData:       make([]byte, BufferSize),
		decryptedData: make([]byte, BufferSize),
		length:        0,
		offset:        0,
	}
}

type DataBuffer struct {
	rawData       []byte
	decryptedData []byte
	length        int
	offset        int
}

func (buffer DataBuffer) printBuffer() {
	if buffer.offset >= buffer.length {
		return
	}
	logging.Debug("Buffer length %d\nraw:\t\t% x\ndecrypted:\t% x\n", buffer.length-buffer.offset, buffer.rawData[buffer.offset:buffer.length], buffer.decryptedData[buffer.offset:buffer.length])
}
