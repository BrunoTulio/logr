module complete-example

go 1.24.0

require github.com/BrunoTulio/logr v0.0.0

require (
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)

replace github.com/BrunoTulio/logr => ../../
