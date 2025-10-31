package zerolog

import (
	"time"

	"github.com/BrunoTulio/logr"
)

func buildAttrs(fields []logr.Field) map[string]any {
	m := make(map[string]any, len(fields))
	for _, f := range fields {
		switch f.Type {
		case logr.StringType:
			m[f.Key] = f.Value.(string)
		case logr.BoolType:
			m[f.Key] = f.Value.(bool)
		case logr.IntType:
			m[f.Key] = f.Value.(int)
		case logr.Uint64Type:
			m[f.Key] = f.Value.(uint64)
		case logr.Float64Type:
			m[f.Key] = f.Value.(float64)
		case logr.TimeType:
			m[f.Key] = f.Value.(time.Time)
		case logr.DurationType:
			m[f.Key] = f.Value.(time.Duration)
		case logr.GroupType:
			groupFields := f.Value.([]logr.Field)
			m[f.Key] = buildAttrs(groupFields)
		default:
			// ignora tipos desconhecidos
		}
	}
	return m
}
