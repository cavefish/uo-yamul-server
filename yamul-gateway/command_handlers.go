package main

import (
	"fmt"
)

type CommandHandler func(client *ClientConnection, commandCode byte)

var clientCommandHandlers = make([]CommandHandler, 256)

func setupCommandHandlers() {
	for i := 0; i < 256; i++ {
		clientCommandHandlers[i] = noop
	}
	clientCommandHandlers[0x80] = loginRequest
	clientCommandHandlers[0x82] = forbiddenClientCommand
	clientCommandHandlers[0xef] = newSeed
}

func noop(client *ClientConnection, commandCode byte) {
	client.err = fmt.Errorf("unknown command %x", commandCode)
}

func forbiddenClientCommand(client *ClientConnection, commandCode byte) {
	client.err = fmt.Errorf("forbidden command %x", commandCode)
}

type LoginRequestCommand struct {
	username string
	password string
	nextkey  byte
}

func loginRequest(client *ClientConnection, commandCode byte) { // 0x80
	username := readFixedString(client, 30)
	password := readFixedString(client, 30)
	nextKey := readByte(client)
	body := LoginRequestCommand{username: username, password: password, nextkey: nextKey}
	go onLoginRequest(client, body)
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

func loginDenied(client *ClientConnection, response LoginDeniedCommand) { // 0x82
	client.outputMutex.Lock()
	defer client.outputMutex.Unlock()
	writeByte(client, 0x82)
	writeByte(client, byte(response.reason))
	_ = client.sendAnyData()
	client.shouldCloseConnection = true
}

type newSeedCommand struct {
	seed            int32
	versionMajor    int32
	versionMinor    int32
	versionRevision int32
	versionPatch    int32
}

func newSeed(client *ClientConnection, commandCode byte) { // 0xef
	seed := readInt(client)
	versionMajor := readInt(client)
	versionMinor := readInt(client)
	versionRevision := readInt(client)
	versionPatch := readInt(client)
	client.encryptSeed = newSeedCommand{seed: seed, versionMajor: versionMajor, versionMinor: versionMinor, versionRevision: versionRevision, versionPatch: versionPatch}
	fmt.Println(client.encryptSeed)
}
