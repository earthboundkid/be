package be

import (
	"testing"
)

// Relaxed returns a testing.TB
// which replaces calls to t.FailNow, t.Fatal, and t.Fatalf
// with calls to t.Fail, t.Error, and t.Errorf respectively.
func Relaxed(t testing.TB) testing.TB {
	return relaxed{t}
}

type relaxed struct {
	testing.TB
}

func (r relaxed) Fatal(args ...any) {
	r.Helper()
	r.Error(args...)
}

func (r relaxed) Fatalf(format string, args ...any) {
	r.Helper()
	r.Errorf(format, args...)
}

func (r relaxed) FailNow() {
	r.Helper()
	r.Fail()
}
