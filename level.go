package logr

type (
	Level int
)

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)
