package be

import (
	"testing"
)

// Relaxed returns a testing.TB which replaces calls to t.Fatalf with calls to t.Errorf.
func Relaxed(t testing.TB) testing.TB {
	return relaxed{t}
}

type relaxed struct {
	testing.TB
}

func (r relaxed) Fatalf(format string, args ...any) {
	r.Helper()
	r.Errorf(format, args...)
}
