package logrus

import (
	"github.com/sirupsen/logrus"
)

func buildLevel(level string) logrus.Level {
	switch level {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
