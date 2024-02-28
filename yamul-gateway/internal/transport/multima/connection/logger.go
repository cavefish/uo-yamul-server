package connection

import (
	log "github.com/sirupsen/logrus"
	"yamul-gateway/internal/interfaces"
)

type logger struct {
	interfaces.Logger
	fields log.Fields
}

func (logger *logger) Errorf(format string, vars ...any) {
	if len(vars) == 0 {
		logger.Error(format)
		return
	}
	log.WithFields(logger.fields).Errorf(format, vars...)
}

func (logger *logger) Error(value ...any) {
	log.WithFields(logger.fields).Error(value...)
}

func (logger *logger) Warningf(format string, vars ...any) {
	if len(vars) == 0 {
		logger.Warning(format)
		return
	}
	log.WithFields(logger.fields).Warnf(format, vars...)
}

func (logger *logger) Warning(value ...any) {
	log.WithFields(logger.fields).Warn(value...)
}

func (logger *logger) Infof(format string, vars ...any) {
	if len(vars) == 0 {
		logger.Info(format)
		return
	}
	log.WithFields(logger.fields).Infof(format, vars...)
}

func (logger *logger) Info(value ...any) {
	log.WithFields(logger.fields).Info(value...)
}

func (logger *logger) Debugf(format string, vars ...any) {
	if len(vars) == 0 {
		logger.Debug(format)
		return
	}
	log.WithFields(logger.fields).Debugf(format, vars...)
}

func (logger *logger) Debug(value ...any) {
	log.WithFields(logger.fields).Debug(value...)
}

func (logger *logger) SetLogField(field string, value any) {
	logger.fields[field] = value
}

func (logger *logger) ClearLogField(field string) {
	delete(logger.fields, field)
}

func CreateAnonymousLogger(name string) interfaces.Logger {
	l := &logger{fields: map[string]interface{}{}}
	l.SetLogField("name", name)
	return l
}

func CreateConnectionLogger(name string, connection *clientConnection) interfaces.Logger {
	l := CreateAnonymousLogger(name)
	l.SetLogField("remote-address", connection.connection.RemoteAddr())
	return l
}
