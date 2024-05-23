package commands

type UpdateObject struct {
	Serial        uint32
	GraphicId     uint16
	XLoc          uint16
	YLoc          uint16
	ZLoc          byte
	Direction     byte
	Hue           uint16
	Flags         byte
	NotorietyFlag byte
}
