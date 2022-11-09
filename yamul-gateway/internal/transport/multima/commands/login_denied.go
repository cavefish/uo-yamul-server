package commands

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
