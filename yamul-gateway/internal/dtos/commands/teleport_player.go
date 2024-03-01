package commands

type TeleportPlayer struct {
	Serial    uint32
	Status    byte
	XLoc      uint16
	YLoc      uint16
	Direction byte
	ZLoc      int8
}
