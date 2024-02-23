package handlers

import (
	"fmt"
	"yamul-gateway/internal/interfaces"
)

func newSeed(client interfaces.ClientConnection) { // 0xef
	seed := client.ReadUInt()
	versionMajor := client.ReadUInt()
	versionMinor := client.ReadUInt()
	versionRevision := client.ReadUInt()
	versionPatch := client.ReadUInt()

	version := fmt.Sprintf("%d.%d.%d.%d.", versionMajor, versionMinor, versionRevision, versionPatch)
	client.GetEncryptionState().Version = version
	client.UpdateEncryptionSeed(seed)
	client.GetLogger().Debug("Encryption reset on login. Version %s", version)
}
