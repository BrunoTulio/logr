package logrus

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/BrunoTulio/logr"
)

type ctxKey struct{}

var _ logr.Logger = (*logger)(nil)

type logger struct {
	logger *logrus.Entry
	writer io.Writer
	fields logr.Fields
	option *Option
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
	args := buildFields(newFields)
	return &logger{
		option: l.option,
		logger: l.logger.WithFields(args),
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

func newLogger(o *Option, fields ...logr.Field) *logger {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.InfoLevel)

	if o.AddSource {
		logrusLogger.SetReportCaller(true)
	}

	var writers []io.Writer

	if o.Console.Enabled {
		logrusLogger.SetLevel(buildLevel(o.Console.Level))
		logrusLogger.SetOutput(os.Stdout)
		logrusLogger.SetFormatter(buildFormatter(o.Console.Formatter))
		writers = append(writers, os.Stdout)
	}

	if o.File.Enabled {
		fileWriter := &lumberjack.Logger{
			Filename: path.Join(o.File.Path, o.File.Name),
			MaxSize:  o.File.MaxSize,
			MaxAge:   o.File.MaxAge,
			Compress: o.File.Compress,
		}
		logrusLogger.AddHook(&WriterHook{
			Writer:    fileWriter,
			Formatter: buildFormatter(o.File.Formatter),
			Level:     buildLevel(o.File.Level),
		})
		writers = append(writers, fileWriter)
	}

	combinedWriter := io.MultiWriter(writers...)

	l := &logger{
		option: o,
		logger: logrus.NewEntry(logrusLogger),
		writer: combinedWriter,
		fields: fields,
	}
	return l
}

func buildFormatter(formatter string) logrus.Formatter {
	switch formatter {
	case "JSON":
		return &logrus.JSONFormatter{}
	default:
		return &logrus.TextFormatter{}
	}
}

func options(fns []FnOption) *Option {
	option := defaultOption()

	for _, fn := range fns {
		fn(option)
	}
	return option
}

// WriterHook is a logrus hook that writes to a custom writer.
type WriterHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
	Level     logrus.Level
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	if entry.Level < hook.Level {
		return nil
	}
	formatted, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(formatted)
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
