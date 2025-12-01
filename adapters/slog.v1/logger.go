package slog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/BrunoTulio/logr"
)

type ctxKey struct{}

var _ logr.Logger = (*logger)(nil)

type logger struct {
	logger  *slog.Logger
	handler slog.Handler
	writer  io.Writer
	fields  logr.Fields
	option  *Option
}

// Info implements logger.Logger.
func (l *logger) Info(message string) {
	l.logger.Info(message)
}

// Infof implements logger.Logger.
func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

// Warn implements logger.Logger.
func (l *logger) Warn(message string) {
	l.logger.Warn(message)
}

// Warnf implements logger.Logger.
func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

// Debug implements logger.Logger.
func (l *logger) Debug(message string) {
	l.logger.Debug(message)
}

// Debugf implements logger.Logger.
func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

// Error implements logger.Logger.
func (l *logger) Error(message string) {
	l.logger.Error(message)
}

// Errorf implements logger.Logger.
func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

// Fatal implements logger.Logger.
func (l *logger) Fatal(message string) {
	l.logger.Error(message)
	os.Exit(1)
}

// Fatalf implements logger.Logger.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}

// FromContext implements logger.Logger.
func (l *logger) FromContext(ctx context.Context) logr.Logger {
	fields, ok := ctx.Value(ctxKey{}).(logr.Fields)
	if !ok {
		fields = logr.Fields{}
	}
	args := buildAttrs(fields)
	return &logger{
		option:  l.option,
		logger:  slog.New(l.handler).With(args...),
		handler: l.handler,
		writer:  l.writer,
		fields:  fields,
	}
}

// GetFields implements logger.Logger.
func (l *logger) GetFields() logr.Fields {
	return l.fields
}

// Output implements logger.Logger.
func (l *logger) Output() io.Writer {
	return l.writer
}

// ToContext implements logger.Logger.
func (l *logger) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, l.fields)
}

// WithField implements logger.Logger.
func (l *logger) WithField(field logr.Field) logr.Logger {
	return l.WithFields(field)
}

// WithFields implements logger.Logger.
func (l *logger) WithFields(fields ...logr.Field) logr.Logger {
	//nolint:gocritic // appendAssign: necessário criar nova slice para manter imutabilidade
	newFields := append(l.fields, fields...)
	args := buildAttrs(newFields)

	return &logger{
		option:  l.option,
		logger:  slog.New(l.handler).With(args...),
		writer:  l.writer,
		fields:  newFields,
		handler: l.handler,
	}
}

// WithMap implements logger.Logger.
func (l *logger) WithMap(m map[string]any) logr.Logger {
	fields := make(logr.Fields, 0, len(m))
	for k, v := range m {
		fields = append(fields, logr.Any(k, v))
	}

	return l.WithFields(fields...)
}

func New(fns ...FnOption) logr.Logger {
	option := options(fns)
	return NewWithOption(option)
}

func NewWithOption(o *Option) logr.Logger {
	l := newLogger(o)
	logr.Set(l)
	return l
}

func newLogger(o *Option, fields ...logr.Field) *logger {
	handler, writer := buildHandlerAndWrite(o)

	l := &logger{
		option:  o,
		logger:  slog.New(handler),
		writer:  writer,
		handler: handler,
		fields:  fields,
	}
	return l
}

func buildHandlerOption(level string, addSource bool) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: addSource,
		Level:     buildLevel(level),
	}
}

func buildFormatter(w io.Writer, formatter string, opts *slog.HandlerOptions) slog.Handler {
	switch formatter {
	case "JSON":
		return slog.NewJSONHandler(w, opts)
	case "TEXT":
		return slog.NewTextHandler(w, opts)
	default:
		return slog.NewTextHandler(w, opts)
	}
}

func buildHandlerAndWrite(o *Option) (slog.Handler, io.Writer) {
	var handlers []slog.Handler
	var writers []io.Writer

	if o.Console.Enabled {
		consoleWriter := os.Stdout
		consoleHandler := buildFormatter(consoleWriter,
			o.Console.Formatter,
			buildHandlerOption(o.Console.Level, o.AddSource),
		)
		handlers = append(handlers, consoleHandler)
		writers = append(writers, consoleWriter)
	}

	if o.File.Enabled {
		fileWriter := &lumberjack.Logger{
			Filename: path.Join(o.File.Path, o.File.Name),
			MaxSize:  o.File.MaxSize,
			MaxAge:   o.File.MaxAge,
			Compress: o.File.Compress,
		}
		consoleHandler := buildFormatter(fileWriter,
			o.File.Formatter,
			buildHandlerOption(o.File.Level, o.AddSource),
		)
		handlers = append(handlers, consoleHandler)
		writers = append(writers, fileWriter)
	}

	if len(handlers) == 0 {
		discardWrite := io.Discard
		discardHandler := slog.NewTextHandler(discardWrite, nil)
		handlers = append(handlers, discardHandler)
		writers = append(writers, discardWrite)
	}

	combinedHandler := NewMultiHandler(handlers...)
	combinedWriter := io.MultiWriter(writers...)
	return combinedHandler, combinedWriter
}

func options(fns []FnOption) *Option {
	option := defaultOption()

	for _, fn := range fns {
		fn(option)
	}
	return option
}
