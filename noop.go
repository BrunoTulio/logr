package logr

import (
	"context"
	"io"
)

var _ Logger = Noop{}

type Noop struct{}

// Debug implements Logger.
func (n Noop) Debug(message string) {}

// Debugf implements Logger.
func (n Noop) Debugf(format string, args ...interface{}) {}

// Error implements Logger.
func (n Noop) Error(message string) {}

// Errorf implements Logger.
func (n Noop) Errorf(format string, args ...interface{}) {}

// Fatal implements Logger.
func (n Noop) Fatal(message string) {}

// Fatalf implements Logger.
func (n Noop) Fatalf(format string, args ...interface{}) {}

// Fields implements Logger.
func (n Noop) GetFields() Fields {
	return Fields{}
}

// FromContext implements Logger.
func (n Noop) FromContext(ctx context.Context) Logger {
	return n
}

// Info implements Logger.
func (n Noop) Info(message string) {}

// Infof implements Logger.
func (n Noop) Infof(format string, args ...interface{}) {}

// Output implements Logger.
func (n Noop) Output() io.Writer {
	return io.Discard
}

// Panic implements Logger.
func (n Noop) Panic(message string) {}

// Panicf implements Logger.
func (n Noop) Panicf(format string, args ...interface{}) {}

// ToContext implements Logger.
func (n Noop) ToContext(ctx context.Context) context.Context {
	return ctx
}

// Warn implements Logger.
func (n Noop) Warn(message string) {}

// Warnf implements Logger.
func (n Noop) Warnf(format string, args ...interface{}) {}

// WithField implements Logger.
func (n Noop) WithField(field Field) Logger {
	return n
}

// WithFields implements Logger.
func (n Noop) WithFields(fields ...Field) Logger {
	return n
}
