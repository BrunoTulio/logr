package zerolog

import (
	"context"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/BrunoTulio/logr"
)

type ctxKey struct{}

const callerSkip = 8

var _ logr.Logger = (*logger)(nil)

type logger struct {
	logger *zerolog.Logger
	writer io.Writer
	fields logr.Fields
	option *Option
}

// Debug implements logr.Logger.
func (l *logger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

// Debugf implements logr.Logger.
func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

// Error implements logr.Logger.
func (l *logger) Error(message string) {
	l.logger.Error().Msg(message)
}

// Errorf implements logr.Logger.
func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

// Fatal implements logr.Logger.
func (l *logger) Fatal(message string) {
	l.logger.Fatal().Msg(message)
}

// Fatalf implements logr.Logger.
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
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
	l.logger.Info().Msg(message)
}

// Infof implements logr.Logger.
func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

// Output implements logr.Logger.
func (l *logger) Output() io.Writer {
	return l.writer
}

// ToContext implements logr.Logger.
func (l *logger) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, l.fields)
}

// Warn implements logr.Logger.
func (l *logger) Warn(message string) {
	l.logger.Warn().Msg(message)
}

// Warnf implements logr.Logger.
func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

// WithField implements logr.Logger.
func (l *logger) WithField(field logr.Field) logr.Logger {
	return l.WithFields(field)
}

// WithFields implements logr.Logger.
func (l *logger) WithFields(fields ...logr.Field) logr.Logger {
	//nolint:gocritic // appendAssign: necessário criar nova slice para manter imutabilidade
	newFields := append(l.fields, fields...)
	args := buildAttrs(newFields)
	newLogger := l.logger.With().Fields(args).Logger()

	return &logger{
		option: l.option,
		logger: &newLogger,
		writer: l.writer,
		fields: newFields,
	}
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
	log, writer := buildLoggerAndWriter(o)
	l := &logger{
		option: o,
		logger: &log,
		writer: writer,
		fields: fields,
	}
	return l
}

func buildLoggerAndWriter(o *Option) (zerolog.Logger, io.Writer) {
	var writers []io.Writer

	// Configura formato de hora padrão
	zerolog.TimeFieldFormat = time.RFC3339

	// Ajusta nível padrão
	level := buildLevel(o.Level)

	// Console (stdout)
	if o.Console.Enabled {
		writers = append(writers, createWriter(os.Stdout, o.Formatter, o.Console.ApplyColor))
	}

	// Arquivo (com rotação via lumberjack)
	if o.File.Enabled {
		fileWriter := &lumberjack.Logger{
			Filename: path.Join(o.File.Path, o.File.Name),
			MaxSize:  o.File.MaxSize,
			MaxAge:   o.File.MaxAge,
			Compress: o.File.Compress,
		}
		writers = append(writers, createWriter(fileWriter, o.Formatter, false))
	}

	if len(writers) == 0 {
		writers = append(writers, io.Discard)
	}

	multi := io.MultiWriter(writers...)

	logger := zerolog.New(multi).
		Level(level).
		With().
		Timestamp().
		CallerWithSkipFrameCount(callerSkip). // se quiser o arquivo/linha
		Logger()

	return logger, multi
}

func createWriter(out io.Writer, formatter string, applyColor bool) io.Writer {
	if formatter == "TEXT" {
		w := zerolog.ConsoleWriter{
			Out:        out,
			TimeFormat: time.RFC3339,
			NoColor:    !applyColor,
		}
		return w
	}
	return out
}

func options(fns []FnOption) *Option {
	option := defaultOption()

	for _, fn := range fns {
		fn(option)
	}
	return option
}
