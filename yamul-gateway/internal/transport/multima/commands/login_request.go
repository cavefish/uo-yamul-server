package commands

type LoginRequestCommand struct {
	Username string
	Password string
	Nextkey  byte
}
