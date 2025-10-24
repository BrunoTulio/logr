package benchmarks

import (
	"testing"

	"github.com/BrunoTulio/logr"
	"github.com/BrunoTulio/logr/adapters/slog.v1"
	"github.com/BrunoTulio/logr/adapters/zap.v1"
)

// Benchmark para comparar performance entre slog e zap.
func BenchmarkSlog_SimpleLog(b *testing.B) {
	logger := slog.New(
		slog.WithConsole(false), // Sem output para benchmark
	)

	b.ResetTimer()
	for range b.N {
		logger.Info("Simple log message")
	}
}

func BenchmarkZap_SimpleLog(b *testing.B) {
	logger := zap.New(
		zap.WithConsole(false), // Sem output para benchmark
	)

	b.ResetTimer()
	for range b.N {
		logger.Info("Simple log message")
	}
}

// Benchmark com campos estruturados.
func BenchmarkSlog_WithFields(b *testing.B) {
	logger := slog.New(
		slog.WithConsole(false),
	)

	b.ResetTimer()
	for i := range b.N {
		logger.WithFields(
			logr.String("user_id", "12345"),
			logr.Bool("active", true),
			logr.Int("count", i),
		).Info("Log with fields")
	}
}

func BenchmarkZap_WithFields(b *testing.B) {
	logger := zap.New(
		zap.WithConsole(false),
	)

	b.ResetTimer()
	for i := range b.N {
		logger.WithFields(
			logr.String("user_id", "12345"),
			logr.Bool("active", true),
			logr.Int("count", i),
		).Info("Log with fields")
	}
}

// Benchmark com campos agrupados.
func BenchmarkSlog_WithGroupedFields(b *testing.B) {
	logger := slog.New(
		slog.WithConsole(false),
	)

	b.ResetTimer()
	for range b.N {
		logger.WithFields(
			logr.Group("user",
				logr.String("id", "12345"),
				logr.String("name", "João"),
				logr.Bool("active", true),
			),
			logr.Group("order",
				logr.String("id", "ORD-123"),
				logr.Float64("total", 99.99),
			),
		).Info("Log with grouped fields")
	}
}

func BenchmarkZap_WithGroupedFields(b *testing.B) {
	logger := zap.New(
		zap.WithConsole(false),
	)

	b.ResetTimer()
	for range b.N {
		logger.WithFields(
			logr.Group("user",
				logr.String("id", "12345"),
				logr.String("name", "João"),
				logr.Bool("active", true),
			),
			logr.Group("order",
				logr.String("id", "ORD-123"),
				logr.Float64("total", 99.99),
			),
		).Info("Log with grouped fields")
	}
}

// Benchmark do logger global.
func BenchmarkGlobalLogger(b *testing.B) {
	logger := zap.New(
		zap.WithConsole(false),
	)
	logr.Set(logger)

	b.ResetTimer()
	for range b.N {
		logr.Info("Global log message")
	}
}

// Benchmark de criação de logger.
func BenchmarkLoggerCreation_Slog(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		_ = slog.New(
			slog.WithConsole(false),
		)
	}
}

func BenchmarkLoggerCreation_Zap(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		_ = zap.New(
			zap.WithConsole(false),
		)
	}
}

// Benchmark de formatação JSON vs TEXT.
func BenchmarkSlog_JSONFormat(b *testing.B) {
	logger := slog.New(
		slog.WithConsole(false),
		slog.WithConsoleFormatter("JSON"),
	)

	b.ResetTimer()
	for i := range b.N {
		logger.WithFields(
			logr.String("message", "test"),
			logr.Int("number", i),
		).Info("JSON formatted log")
	}
}

func BenchmarkSlog_TEXTFormat(b *testing.B) {
	logger := slog.New(
		slog.WithConsole(false),
		slog.WithConsoleFormatter("TEXT"),
	)

	b.ResetTimer()
	for i := range b.N {
		logger.WithFields(
			logr.String("message", "test"),
			logr.Int("number", i),
		).Info("TEXT formatted log")
	}
}
