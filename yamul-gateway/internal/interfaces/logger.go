package interfaces

type Logger interface {
	Error(value ...any)
	Errorf(format string, vars ...any)
	Warning(value ...any)
	Warningf(format string, vars ...any)
	Info(value ...any)
	Infof(format string, vars ...any)
	Debug(value ...any)
	Debugf(format string, vars ...any)
	SetLogField(field string, value any)
	ClearLogField(field string)
}
