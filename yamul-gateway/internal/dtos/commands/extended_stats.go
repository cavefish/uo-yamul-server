package commands

const (
	ExtendedStats_Dead          = 0
	ExtendedStats_AttributeLock = 2
)

type ExtendedStats struct {
	Type     byte
	ObjectID uint32
	StrLock  bool
	DexLock  bool
	IntLock  bool
	IsDead   bool
}
