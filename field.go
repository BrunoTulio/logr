package logr

import "time"

type (
	FieldType int

	Fields []Field

	Field struct {
		Type  FieldType
		Key   string
		Value any
	}
)

const (
	StringType FieldType = iota
	BoolType
	IntType
	Uint64Type
	Float64Type
	TimeType
	DurationType
	GroupType
)

func String(key, value string) Field {
	return Field{Key: key, Value: value, Type: StringType}
}

func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value, Type: BoolType}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value, Type: IntType}
}

func Uint64(key string, value uint64) Field {
	return Field{Key: key, Value: value, Type: Uint64Type}
}

func Float64(key string, value float64) Field {
	return Field{Key: key, Value: value, Type: Float64Type}
}

func Time(key string, value time.Time) Field {
	return Field{Key: key, Value: value, Type: TimeType}
}

func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Value: value, Type: DurationType}
}

func Group(name string, fields ...Field) Field {
	return Field{Key: name, Value: fields, Type: GroupType}
}
