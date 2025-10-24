package slog

import (
	"log/slog"
	"time"

	"github.com/BrunoTulio/logr"
)

func buildAttrGroup(fields []logr.Field) []any {
	result := make([]any, 0, len(fields))
	for _, f := range fields {
		result = append(result, buildAttr(f))
	}
	return result
}

func buildAttr(f logr.Field) slog.Attr {
	switch f.Type {
	case logr.StringType:
		return slog.String(f.Key, f.Value.(string))
	case logr.BoolType:
		return slog.Bool(f.Key, f.Value.(bool))
	case logr.IntType:
		return slog.Int(f.Key, f.Value.(int))
	case logr.Uint64Type:
		return slog.Uint64(f.Key, f.Value.(uint64))
	case logr.Float64Type:
		return slog.Float64(f.Key, f.Value.(float64))
	case logr.TimeType:
		return slog.Time(f.Key, f.Value.(time.Time))
	case logr.DurationType:
		return slog.Duration(f.Key, f.Value.(time.Duration))
	case logr.GroupType:
		groupFields := f.Value.([]logr.Field)
		return slog.Group(f.Key, buildAttrGroup(groupFields)...)
	default:
		return slog.Attr{}
	}
}

func buildAttrs(fields logr.Fields) []any {
	result := make([]any, 0, len(fields))
	for _, f := range fields {
		result = append(result, buildAttr(f))
	}
	return result
}
