package login

import (
	"strings"
)

const (
	OK int = iota
	INVALID_USER
	INVALID_CREDENTIALS
)

func CheckUserCredentials(username string, password string) (bool, int) {
	if !strings.EqualFold(username, "admin") {
		return false, INVALID_USER
	}
	if !strings.EqualFold(password, "admin") {
		return false, INVALID_CREDENTIALS
	}
	return true, OK

}
