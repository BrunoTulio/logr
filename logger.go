package logr

import (
	"context"
	"io"
)

type Logger interface {
	Info(message string)
	Infof(format string, args ...interface{})

	Warn(message string)
	Warnf(format string, args ...interface{})

	Error(message string)
	Errorf(format string, args ...interface{})

	Fatal(message string)
	Fatalf(format string, args ...interface{})

	Debug(message string)
	Debugf(format string, args ...interface{})

	WithFields(fields ...Field) Logger
	WithField(field Field) Logger

	ToContext(ctx context.Context) context.Context
	FromContext(ctx context.Context) Logger
	GetFields() Fields

	Output() io.Writer
}
