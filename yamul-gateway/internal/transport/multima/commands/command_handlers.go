package commands

import (
	"fmt"
	"yamul-gateway/internal/transport/multima/connection"
)

func SetupCommandHandlers() {
	for i := 0; i < 256; i++ {
		connection.ClientCommandHandlers[i] = noop
	}
	connection.ClientCommandHandlers[0x80] = loginRequest
	connection.ClientCommandHandlers[0x82] = forbiddenClientCommand
	connection.ClientCommandHandlers[0xef] = newSeed
}

func noop(client *connection.ClientConnection, commandCode byte) {
	client.Err = fmt.Errorf("unknown command %x", commandCode)
}

func forbiddenClientCommand(client *connection.ClientConnection, commandCode byte) {
	client.Err = fmt.Errorf("forbidden command %x", commandCode)
}

type LoginRequestCommand struct {
	username string
	password string
	nextkey  byte
}

func loginRequest(client *connection.ClientConnection, commandCode byte) { // 0x80
	username := client.ReadFixedString(30)
	password := client.ReadFixedString(30)
	nextKey := client.ReadByte()
	body := LoginRequestCommand{username: username, password: password, nextkey: nextKey}
	go OnLoginRequest(client, body)
}

type LoginDeniedReason byte

const (
	incorrectUsernamePassword LoginDeniedReason = iota
	accountAlreadyInUse
	accountBlocked
	badPassword
	communicationProblem
	igrConcurrencyLimit
	igrTimeLimit
	igrGeneralFailure
)

type LoginDeniedCommand struct {
	reason LoginDeniedReason
}

func loginDenied(client *connection.ClientConnection, response LoginDeniedCommand) { // 0x82
	client.Lock()
	defer client.Unlock()
	client.WriteByte(0x82)
	client.WriteByte(byte(response.reason))
	_ = client.SendAnyData()
	client.ShouldCloseConnection = true
}

type NewSeedCommand struct {
	seed            int32
	versionMajor    int32
	versionMinor    int32
	versionRevision int32
	versionPatch    int32
}

func newSeed(client *connection.ClientConnection, commandCode byte) { // 0xef
	seed := client.ReadInt()
	versionMajor := client.ReadInt()
	versionMinor := client.ReadInt()
	versionRevision := client.ReadInt()
	versionPatch := client.ReadInt()
	client.EncryptSeed = connection.EncryptionConfig{Seed: seed, VersionMajor: versionMajor, VersionMinor: versionMinor, VersionRevision: versionRevision, VersionPatch: versionPatch}
	fmt.Println(client.EncryptSeed)
}
