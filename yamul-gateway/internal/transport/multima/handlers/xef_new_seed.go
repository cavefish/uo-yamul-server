package handlers

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

func newSeed(client *connection.ClientConnection) { // 0xef
	seed := client.ReadUInt()
	client.EncryptionState.VersionMajor = client.ReadUInt()
	client.EncryptionState.VersionMinor = client.ReadUInt()
	client.EncryptionState.VersionRevision = client.ReadUInt()
	client.EncryptionState.VersionPatch = client.ReadUInt()

	client.UpdateEncryptionSeed(seed)
	logging.Debug("Encryption seed %+v\n", client.EncryptionState)
}
