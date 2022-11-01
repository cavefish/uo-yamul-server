package main

import "fmt"

func readByte(client *ClientConnection) byte {
	if client.receiveData() != nil {
		return 0
	}
	value := client.inputBuffer.decryptedData[client.inputBuffer.offset]
	client.inputBuffer.offset++
	return value
}

func writeByte(client *ClientConnection, value byte) {
	if client.sendDataIfAlmostFull(1) != nil {
		return
	}
	client.outputBuffer.decryptedData[client.outputBuffer.length] = value
	client.outputBuffer.length++
}

func readInt(client *ClientConnection) int32 {
	value := int32(readByte(client))
	value = value<<8 | int32(readByte(client))
	value = value<<8 | int32(readByte(client))
	value = value<<8 | int32(readByte(client))
	return value
}

func readFixedString(client *ClientConnection, length int) string {
	value := make([]byte, length)
	for i := 0; i < length; i++ {
		value[i] = readByte(client)
	}
	return string(value)
}

func printBuffer(buffer DataBuffer) {
	fmt.Printf("Buffer length %d\nraw:\t\t% x\n", buffer.length, buffer.rawData[0:buffer.length])
	fmt.Printf("decrypted:\t% x\n", buffer.decryptedData[0:buffer.length])
}
