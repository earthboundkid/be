package be_test

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func runTest(test func(*testing.T)) bool {
	return testing.RunTests(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{{"test", test}})
}

func TestRelaxed(t *testing.T) {
	finished := false
	be.False(t, runTest(func(t *testing.T) {
		rt := be.Relaxed(t)
		rt.FailNow()
		rt.Fatal("boom!")
		rt.Fatalf("msg: %v", "boom!")
		finished = true
	}))
	be.True(t, finished)
}
