package commands

type UnicodeSpeechSelected struct {
	Mode     byte
	Hue      uint16
	Font     uint16
	Language string
	Keywords []uint16
	Text     string
}
