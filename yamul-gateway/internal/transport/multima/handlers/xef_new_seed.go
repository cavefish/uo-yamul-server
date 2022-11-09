package handlers

import (
	"yamul-gateway/internal/logging"
	"yamul-gateway/internal/transport/multima/connection"
)

func newSeed(client *connection.ClientConnection) { // 0xef
	seed := client.ReadUInt()
	versionMajor := client.ReadUInt()
	versionMinor := client.ReadUInt()
	versionRevision := client.ReadUInt()
	versionPatch := client.ReadUInt()
	client.EncryptSeed = connection.EncryptionConfig{Seed: seed, VersionMajor: versionMajor, VersionMinor: versionMinor, VersionRevision: versionRevision, VersionPatch: versionPatch}
	logging.Debug("Encryption seed %+v\n", client.EncryptSeed)
}
