package xef_newSeed

import (
	"fmt"
	"yamul-gateway/internal/transport/multima/connection"
)

type NewSeedCommand struct {
	seed            int32
	versionMajor    int32
	versionMinor    int32
	versionRevision int32
	versionPatch    int32
}

func NewSeed(client *connection.ClientConnection, commandCode byte) { // 0xef
	seed := client.ReadInt()
	versionMajor := client.ReadInt()
	versionMinor := client.ReadInt()
	versionRevision := client.ReadInt()
	versionPatch := client.ReadInt()
	client.EncryptSeed = connection.EncryptionConfig{Seed: seed, VersionMajor: versionMajor, VersionMinor: versionMinor, VersionRevision: versionRevision, VersionPatch: versionPatch}
	fmt.Println(client.EncryptSeed)
}
