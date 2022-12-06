package commands

type PlayerStartConfirmation struct {
	CharacterID       uint32
	CharacterBodyType uint16
	Coordinates       Coordinates
	DirectionFacing   DirectionFacing
}
