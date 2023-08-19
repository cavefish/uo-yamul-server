package commands

type ClientFeatures struct {
	Unknown0001            bool // 0x0001
	OverwriteConfigButtons bool // 0x0002
	SingleCharacterSlot    bool // 0x0004
	ContextMenus           bool // 0x0008
	LimitCharacterSlots    bool // 0x0010
	EnableAOS              bool // 0x0020
	SixthSlot              bool // 0x0040
	SamuraiNinja           bool // 0x0080
	ElvenRace              bool // 0x0100
	Unknown0200            bool // 0x0200
	UO3D                   bool // 0x0400
	Unknown0800            bool // 0x0800
	SeventhSlot            bool // 0x1000
	Unknown2000            bool // 0x2000
	NewMovement            bool // 0x4000
	UnlockPvPAreas         bool // 0x8000
}
