package connection

import (
	"fmt"
	"time"
	"yamul-gateway/internal/interfaces"
)

const (
	LogLevelError = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

type logger struct {
	interfaces.Logger
	client   *clientConnection
	name     string
	prefix   string
	logLevel int
}

func (logger *logger) Error(format string, vars ...any) {
	logger.log(LogLevelError, format, vars)
}

func (logger *logger) Warning(format string, vars ...any) {
	logger.log(LogLevelWarning, format, vars)
}

func (logger *logger) Info(format string, vars ...any) {
	logger.log(LogLevelInfo, format, vars)
}

func (logger *logger) Debug(format string, vars ...any) {
	logger.log(LogLevelDebug, format, vars)
}

func (logger *logger) SetPrefix(prefix string) {
	logger.prefix = prefix
}

func (logger *logger) log(level int, format string, vars []any) {
	if level > logger.logLevel {
		return
	}

	output := format
	if len(vars) > 0 {
		output = fmt.Sprintf(format, vars...)
	}

	clientPrefix := logger.getClientPrefix()
	levelPrefix := logger.getLogLevelPrefix()

	time := time.Now().Format(time.Stamp)

	fmt.Printf("%s\t[%s]%s%s\t%s\n", time, levelPrefix, clientPrefix, logger.prefix, output)
}

func (logger *logger) getClientPrefix() string {
	if logger.client == nil {
		return logger.name
	}

	return fmt.Sprintf("[%s]", logger.client.connection.RemoteAddr())
}

func (logger *logger) getLogLevelPrefix() string {
	switch logger.logLevel {
	case LogLevelError:
		return "ERROR"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelInfo:
		return "INFO"
	case LogLevelDebug:
		return "DEBUG"
	default:
		return fmt.Sprintf("ERROR-UNKNOWN-LEVEL-%d", logger.logLevel)
	}
}

func LoggerFor(name string) interfaces.Logger {
	return &logger{
		client:   nil,
		name:     name,
		logLevel: LogLevelDebug,
	}
}
