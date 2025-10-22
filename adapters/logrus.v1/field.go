package logrus

import (
	"github.com/BrunoTulio/logr"
	"github.com/sirupsen/logrus"
)

func buildFields(fields logr.Fields) logrus.Fields {
	result := make(logrus.Fields, len(fields))
	for _, f := range fields {
		result[f.Key] = f.Value
	}
	return result
}
