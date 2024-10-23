package commands

type MoveReject struct {
	Sequence  byte
	XLoc      uint16
	YLoc      uint16
	ZLoc      byte
	Direction byte
}
