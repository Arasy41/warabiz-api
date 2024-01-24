package logger

import "github.com/sirupsen/logrus"

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Trace(args ...interface{})
	Tracef(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields logrus.Fields) Logger
}

func (l apiLogger) WithField(key string, value interface{}) Logger {
	l.entry = l.logger.WithField(key, value)
	return &l
}

func (l apiLogger) WithFields(fields logrus.Fields) Logger {
	l.entry = l.entry.WithFields(fields)
	return &l
}

func (l *apiLogger) Debug(args ...interface{}) {
	l.getEntry().Debug(args...)
}

func (l *apiLogger) Debugf(template string, args ...interface{}) {
	l.getEntry().Debugf(template, args...)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.getEntry().Info(args...)
}

func (l *apiLogger) Infof(template string, args ...interface{}) {
	l.getEntry().Infof(template, args...)
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.getEntry().Warn(args...)
}

func (l *apiLogger) Warnf(template string, args ...interface{}) {
	l.getEntry().Warnf(template, args...)
}

func (l *apiLogger) Error(args ...interface{}) {
	l.getEntry().Error(args...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	l.getEntry().Errorf(template, args...)
}

func (l *apiLogger) Trace(args ...interface{}) {
	l.getEntry().Trace(args...)
}

func (l *apiLogger) Tracef(template string, args ...interface{}) {
	l.getEntry().Tracef(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.getEntry().Panic(args...)
}

func (l *apiLogger) Panicf(template string, args ...interface{}) {
	l.getEntry().Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.getEntry().Fatal(args...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.getEntry().Fatalf(template, args...)
}
