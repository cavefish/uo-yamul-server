package connection

import (
	"yamul-gateway/internal/logging"
)

type DataBuffer struct {
	rawData       []byte
	decryptedData []byte
	length        int `default:"0"`
	offset        int `default:"0"`
}

func (buffer DataBuffer) printBuffer() {
	if buffer.offset >= buffer.length {
		return
	}
	logging.Debug("Buffer length %d\nraw:\t\t% x\ndecrypted:\t% x\n", buffer.length-buffer.offset, buffer.rawData[buffer.offset:buffer.length], buffer.decryptedData[buffer.offset:buffer.length])
}
