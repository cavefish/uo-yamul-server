package handlers

import (
	"fmt"
	"yamul-gateway/internal/transport/multima/connection"
)

func newSeed(client *connection.ClientConnection) { // 0xef
	seed := client.ReadUInt()
	versionMajor := client.ReadUInt()
	versionMinor := client.ReadUInt()
	versionRevision := client.ReadUInt()
	versionPatch := client.ReadUInt()

	version := fmt.Sprintf("%d.%d.%d.%d.", versionMajor, versionMinor, versionRevision, versionPatch)
	client.EncryptionState.Version = version
	client.UpdateEncryptionSeed(seed)
	client.Logger.Debug("Encryption reset on login. Version %s", version)
}
