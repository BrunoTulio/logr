package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/BrunoTulio/logr"
)

func buildFields(fields logr.Fields) logrus.Fields {
	result := make(logrus.Fields, len(fields))
	for _, f := range fields {
		if f.Type == logr.GroupType {
			groupFields := f.Value.([]logr.Field)
			result[f.Key] = buildFields(groupFields)
		} else {
			result[f.Key] = f.Value
		}
	}
	return result
}
