package commands

type ClientMoveRequest struct {
	Direction byte
	Sequence  byte
	AckKey    uint32
}
