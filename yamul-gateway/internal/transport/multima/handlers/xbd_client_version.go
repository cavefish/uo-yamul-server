package handlers

import (
	"yamul-gateway/internal/interfaces"
)

func receiveClientVersion(client interfaces.ClientConnection) { // 0xBD
	size := client.ReadUShort() - 3
	body := client.ReadFixedBytes(int(size))
	client.GetLogger().Debugf("User logged with version '%s'", body[0:size-1])
}
