package commands

type ClientFeatures struct {
	Unknown0001            bool // 0x0001
	OverwriteConfigButtons bool // 0x0002
	SingleCharacterSlot    bool // 0x0004
	ContextMenus           bool // 0x0008
	LimitCharacterSlots    bool // 0x0010
	EnableAOS              bool // 0x0020
	SixthSlot              bool // 0x0040
}
