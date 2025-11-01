package logrus

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestNewLogger(t *testing.T) {
}
