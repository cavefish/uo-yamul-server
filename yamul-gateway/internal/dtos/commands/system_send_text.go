package commands

type SystemSendText struct {
	Serial uint32
	Model  uint16
	Type   byte
	Hue    uint16
	Font   uint16
	Name   string
	Body   string
}
