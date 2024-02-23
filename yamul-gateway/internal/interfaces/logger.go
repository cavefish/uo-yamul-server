package interfaces

type Logger interface {
	Error(format string, vars ...any)
	Warning(format string, vars ...any)
	Info(format string, vars ...any)
	Debug(format string, vars ...any)
	SetPrefix(prefix string)
}
