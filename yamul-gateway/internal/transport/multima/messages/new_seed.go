package messages

type NewSeedCommand struct {
	seed            int32
	versionMajor    int32
	versionMinor    int32
	versionRevision int32
	versionPatch    int32
}
