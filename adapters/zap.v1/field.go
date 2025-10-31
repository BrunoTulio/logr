package zap

import (
	"github.com/BrunoTulio/logr"
)

const fieldsValuePair = 2

func buildSugaredArgs(fields logr.Fields) []interface{} {
	args := make([]interface{}, 0, len(fields)*fieldsValuePair)
	for _, f := range fields {
		if f.Type == logr.GroupType {
			groupFields := f.Value.([]logr.Field)
			args = append(args, f.Key, buildGroupMap(groupFields))
		} else {
			args = append(args, f.Key, f.Value)
		}
	}
	return args
}

func buildGroupMap(fields logr.Fields) map[string]interface{} {
	m := make(map[string]interface{}, len(fields))
	for _, f := range fields {
		switch f.Type {
		case logr.GroupType:
			m[f.Key] = buildGroupMap(f.Value.([]logr.Field))
		case logr.StringType, logr.BoolType, logr.IntType, logr.Uint64Type, logr.Float64Type, logr.TimeType, logr.DurationType:
			m[f.Key] = f.Value
		}
	}
	return m
}
