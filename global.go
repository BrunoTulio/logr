package logr

import (
	"context"
	"io"
)

var l Logger = Noop{}

func Set(logger Logger) {
	l = logger
}

func Info(message string) {
	l.Info(message)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Warn(message string) {
	l.Warn(message)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Error(message string) {
	l.Error(message)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Fatal(message string) {
	l.Fatal(message)
}

func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func Debug(message string) {
	l.Debug(message)
}

func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func WithFields(field ...Field) Logger {
	return l.WithFields(field...)
}

func WithField(field Field) Logger {
	return l.WithField(field)
}

func ToContext(ctx context.Context) context.Context {
	return l.ToContext(ctx)
}

func FromContext(ctx context.Context) Logger {
	return l.FromContext(ctx)
}

func GetFields() Fields {
	return l.GetFields()
}

func Output() io.Writer {
	return l.Output()
}
