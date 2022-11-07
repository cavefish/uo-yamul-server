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
	logging.Debug("Buffer length %d\nraw:\t\t% x\ndecrypted:\t% x\n", buffer.length, buffer.rawData[0:buffer.length], buffer.decryptedData[0:buffer.length])
}
