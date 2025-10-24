package zap

import (
	"context"
	"io"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/BrunoTulio/logr"
)

type ctxKey struct{}

const callerSkip = 1

var _ logr.Logger = (*logger)(nil)

type logger struct {
	logger *zap.SugaredLogger
	writer io.Writer
	fields logr.Fields
	option *Option
}

// Debug implements logr.Logger.
func (l *logger) Debug(message string) {
	l.logger.Debug(message)
}

// Debugf implements logr.Logger.
func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Error implements logr.Logger.
func (l *logger) Error(message string) {
	l.logger.Error(message)
}

// Errorf implements logr.Logger.
func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal implements logr.Logger.
func (l *logger) Fatal(message string) {
	l.logger.Fatal(message)
}

// Fatalf implements logr.Logger.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// FromContext implements logr.Logger.
func (l *logger) FromContext(ctx context.Context) logr.Logger {
	fields, ok := ctx.Value(ctxKey{}).(logr.Fields)
	if !ok {
		fields = logr.Fields{}
	}
	return l.WithFields(fields...)
}

// GetFields implements logr.Logger.
func (l *logger) GetFields() logr.Fields {
	return l.fields
}

// Info implements logr.Logger.
func (l *logger) Info(message string) {
	l.logger.Info(message)
}

// Infof implements logr.Logger.
func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn implements logr.Logger.
func (l *logger) Warn(message string) {
	l.logger.Warn(message)
}

// Warnf implements logr.Logger.
func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Output implements logr.Logger.
func (l *logger) Output() io.Writer {
	return l.writer
}

// ToContext implements logr.Logger.
func (l *logger) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, l.fields)
}

// WithField implements logr.Logger.
func (l *logger) WithField(field logr.Field) logr.Logger {
	return l.WithFields(field)
}

// WithFields implements logr.Logger.
func (l *logger) WithFields(fields ...logr.Field) logr.Logger {
	//nolint:gocritic // appendAssign: necessário criar nova slice para manter imutabilidade
	newFields := append(l.fields, fields...)
	args := buildSugaredArgs(newFields)

	return &logger{
		option: l.option,
		logger: l.logger.With(args...),
		writer: l.writer,
		fields: newFields,
	}
}

func New(fns ...FnOption) *logger {
	option := options(fns)
	return NewWithOption(option)
}

func NewWithOption(o *Option) *logger {
	l := newLogger(o)
	logr.Set(l)
	return l
}

func buildCoreAndWriter(o *Option) (zapcore.Core, io.Writer) {
	cores := []zapcore.Core{}
	var writers []io.Writer

	if o.Console.Enabled {
		level := buildLevel(o.Console.Level)
		writer := zapcore.Lock(os.Stdout)
		coreconsole := zapcore.NewCore(buildEncoder(o.Console.Formatter), writer, level)
		cores = append(cores, coreconsole)
		writers = append(writers, writer)
	}

	if o.File.Enabled {
		lumber := &lumberjack.Logger{
			Filename: path.Join(o.File.Path, o.File.Name),
			MaxSize:  o.File.MaxSize,
			Compress: o.File.Compress,
			MaxAge:   o.File.MaxAge,
		}

		level := buildLevel(o.File.Level)
		writer := zapcore.AddSync(lumber)
		corefile := zapcore.NewCore(buildEncoder(o.File.Formatter), writer, level)
		cores = append(cores, corefile)
		writers = append(writers, lumber)
	}

	combinedCore := zapcore.NewTee(cores...)
	combinedWriter := io.MultiWriter(writers...)

	return combinedCore, combinedWriter
}

func newLogger(o *Option, fields ...logr.Field) *logger {
	core, writer := buildCoreAndWriter(o)

	sugar := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(callerSkip),
	).Sugar()

	l := &logger{
		option: o,
		logger: sugar,
		writer: writer,
		fields: fields,
	}
	return l
}

func options(fns []FnOption) *Option {
	option := defaultOption()

	for _, fn := range fns {
		fn(option)
	}
	return option
}
