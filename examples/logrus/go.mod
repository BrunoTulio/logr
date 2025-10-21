module logrus-example

go 1.22.10

toolchain go1.24.5

replace github.com/BrunoTulio/logr => ../..

require github.com/BrunoTulio/logr v0.0.0-00010101000000-000000000000

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
