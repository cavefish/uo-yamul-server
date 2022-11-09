package handlers

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

func newSeed(client *connection.ClientConnection) { // 0xef
	seed := client.ReadInt()
	versionMajor := client.ReadInt()
	versionMinor := client.ReadInt()
	versionRevision := client.ReadInt()
	versionPatch := client.ReadInt()
	client.EncryptSeed = connection.EncryptionConfig{Seed: seed, VersionMajor: versionMajor, VersionMinor: versionMinor, VersionRevision: versionRevision, VersionPatch: versionPatch}
	logging.Debug("Encryption seed %+v\n", client.EncryptSeed)
}
