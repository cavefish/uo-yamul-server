package connection

import "fmt"

const BufferSize = 15360

func CreateInputDataBuffer() InputDataBuffer {
	return InputDataBuffer{
		incomingTcpData: make([]byte, BufferSize),
		decryptedData:   make([]byte, BufferSize),
		length:          0,
		offset:          0,
	}
}

func CreateOutputDataBuffer() OutputDataBuffer {
	return OutputDataBuffer{
		outgoingTcpData: make([]byte, BufferSize),
		decryptedData:   make([]byte, BufferSize),
		compressedData:  make([]byte, BufferSize),
		length:          0,
	}
}

type OutputDataBuffer struct {
	outgoingTcpData []byte
	decryptedData   []byte
	compressedData  []byte
	length          int
}

type InputDataBuffer struct {
	incomingTcpData []byte
	decryptedData   []byte
	length          int
	offset          int
}

func (buffer InputDataBuffer) printBuffer() string {
	if buffer.offset >= buffer.length {
		return "Empty buffer"
	}
	return fmt.Sprintf("Input Buffer length %d\nraw:\t\t% x", buffer.length-buffer.offset, buffer.decryptedData[buffer.offset:buffer.length])
}

func (buffer OutputDataBuffer) printBuffer() string {
	return fmt.Sprintf("Output Buffer length %d\nraw:\t\t% x", buffer.length, buffer.decryptedData[:buffer.length])
}
