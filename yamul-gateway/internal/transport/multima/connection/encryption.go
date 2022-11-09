package connection

type EncryptionConfig struct {
	Seed            uint32
	VersionMajor    uint32
	VersionMinor    uint32
	VersionRevision uint32
	VersionPatch    uint32
}
