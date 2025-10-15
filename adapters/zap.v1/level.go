package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildLevel(level string) zapcore.Level {
	switch level {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "WARN":
		return zap.WarnLevel
	case "ERROR":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
