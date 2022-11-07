package x82_loginDenied

import "yamul-gateway/internal/transport/multima/connection"

type LoginDeniedReason byte

const (
	IncorrectUsernamePassword LoginDeniedReason = iota
	AccountAlreadyInUse
	AccountBlocked
	BadPassword
	CommunicationProblem
	IgrConcurrencyLimit
	IgrTimeLimit
	IgrGeneralFailure
)

type LoginDeniedCommand struct {
	Reason LoginDeniedReason
}

func LoginDenied(client *connection.ClientConnection, response LoginDeniedCommand) { // 0x82
	client.Lock()
	defer client.Unlock()
	client.WriteByte(0x82)
	client.WriteByte(byte(response.Reason))
	_ = client.SendAnyData()
	client.ShouldCloseConnection = true
}
