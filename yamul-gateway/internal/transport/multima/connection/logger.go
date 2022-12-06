package connection

import "fmt"

const (
	LOG_LEVEL_ERROR = iota
	LOG_LEVEL_WARNING
	LOG_LEVEL_INFO
	LOG_LEVEL_DEBUG
)

type logger struct {
	Logger
	client   *ClientConnection
	name     string
	logLevel int
}

type Logger interface {
	Error(format string, vars ...any)
	Warning(format string, vars ...any)
	Info(format string, vars ...any)
	Debug(format string, vars ...any)
}

func (logger *logger) Error(format string, vars ...any) {
	logger.log(LOG_LEVEL_ERROR, format, vars)
}

func (logger *logger) Warning(format string, vars ...any) {
	logger.log(LOG_LEVEL_WARNING, format, vars)
}

func (logger *logger) Info(format string, vars ...any) {
	logger.log(LOG_LEVEL_INFO, format, vars)
}

func (logger *logger) Debug(format string, vars ...any) {
	logger.log(LOG_LEVEL_DEBUG, format, vars)
}

func (logger *logger) log(level int, format string, vars ...any) {
	if level > logger.logLevel {
		return
	}

	output := format
	if len(vars) > 0 {
		output = fmt.Sprintf(format, vars)
	}

	clientPrefix := logger.getClientPrefix()
	levelPrefix := logger.getLogLevelPrefix()

	fmt.Printf("[%s] %s\t%s\n", levelPrefix, clientPrefix, output)
}

func (logger *logger) getClientPrefix() string {
	if logger.client == nil {
		return logger.name
	}

	return fmt.Sprintf("[server=%s, client=%s]", logger.client.Connection.LocalAddr(), logger.client.Connection.RemoteAddr())
}

func (logger *logger) getLogLevelPrefix() string {
	switch logger.logLevel {
	case LOG_LEVEL_ERROR:
		return "ERROR"
	case LOG_LEVEL_WARNING:
		return "WARNING"
	case LOG_LEVEL_INFO:
		return "INFO"
	case LOG_LEVEL_DEBUG:
		return "DEBUG"
	default:
		return fmt.Sprintf("ERROR-UNKNOWN-LEVEL-%d", logger.logLevel)
	}
}

func LoggerFor(name string) Logger {
	return &logger{
		client:   nil,
		name:     name,
		logLevel: LOG_LEVEL_DEBUG,
	}
}
