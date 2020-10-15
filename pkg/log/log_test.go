package log

import (
	"testing"
)

func TestLogLevel(t *testing.T) {
	level := "info"
	SetLevel(level)
	if GetLevel() != level {
		t.Fatalf("Expected %q, got %q", level, GetLevel())
	}
}
