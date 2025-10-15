package zap

import (
	"time"

	"github.com/BrunoTulio/logr"

	"go.uber.org/zap/zapcore"
)

func buildSugaredArgs(fields logr.Fields) []interface{} {
	args := make([]interface{}, 0, len(fields)*2)
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
			// recursão para grupos aninhados
			m[f.Key] = buildGroupMap(f.Value.([]logr.Field))
		default:
			m[f.Key] = f.Value
		}
	}
	return m
}

// buildZapFieldToEncoder adiciona um campo logr.Field diretamente no zapcore.ObjectEncoder
func buildZapFieldToEncoder(enc zapcore.ObjectEncoder, f logr.Field) {
	switch f.Type {
	case logr.StringType:
		enc.AddString(f.Key, f.Value.(string))
	case logr.BoolType:
		enc.AddBool(f.Key, f.Value.(bool))
	case logr.IntType:
		enc.AddInt(f.Key, f.Value.(int))
	case logr.Uint64Type:
		enc.AddUint64(f.Key, f.Value.(uint64))
	case logr.Float64Type:
		enc.AddFloat64(f.Key, f.Value.(float64))
	case logr.TimeType:
		enc.AddTime(f.Key, f.Value.(time.Time))
	case logr.DurationType:
		enc.AddDuration(f.Key, f.Value.(time.Duration))
	case logr.GroupType:
		groupFields := f.Value.([]logr.Field)
		enc.AddObject(f.Key, zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			for _, gf := range groupFields {
				buildZapFieldToEncoder(enc, gf)
			}
			return nil
		}))
	default:
	}
}
