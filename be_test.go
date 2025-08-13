package be_test

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"strings"
	"testing"
	"time"

	"github.com/carlmjohnson/be"
)

type testingTB struct {
	testing.TB
	failed bool
	w      io.Writer
}

func (t *testingTB) Helper() {}

func (t *testingTB) Fatalf(format string, args ...any) {
	t.failed = true
	fmt.Fprintf(t.w, format, args...)
}

func Test(t *testing.T) {
	beOkay := func(callback func(tb testing.TB)) {
		t.Helper()
		var buf strings.Builder
		tb := &testingTB{w: &buf}
		callback(tb)
		if tb.failed {
			t.Fatal("failed too soon")
		}
		if buf.String() != "" {
			t.Fatal("wrote too much")
		}
	}
	beOkay(func(tb testing.TB) { be.Zero(tb, time.Time{}.Local()) })
	beOkay(func(tb testing.TB) { be.Zero(tb, []string(nil)) })
	beOkay(func(tb testing.TB) { be.Nonzero(tb, []string{""}) })
	beOkay(func(tb testing.TB) { be.NilErr(tb, nil) })
	beOkay(func(tb testing.TB) { be.True(tb, true) })
	beOkay(func(tb testing.TB) { be.False(tb, false) })
	beOkay(func(tb testing.TB) { be.EqualLength(tb, 0, map[int]int{}) })
	beOkay(func(tb testing.TB) { be.EqualLength(tb, 1, map[int]int{1: 1}) })
	ch := make(chan int, 1)
	beOkay(func(tb testing.TB) { be.EqualLength(tb, 0, ch) })
	ch <- 1
	beOkay(func(tb testing.TB) { be.EqualLength(tb, 1, ch) })
	seq2 := maps.All(map[int]int{1: 1})
	beOkay(func(tb testing.TB) { be.EqualLength(tb, 1, seq2) })
	beBad := func(callback func(tb testing.TB)) {
		t.Helper()
		var buf strings.Builder
		tb := &testingTB{w: &buf}
		callback(tb)
		if !tb.failed {
			t.Fatal("did not fail")
		}
		if buf.String() == "" {
			t.Fatal("wrote too little")
		}
	}
	beBad(func(tb testing.TB) { be.AllEqual(tb, []string{}, []string{""}) })
	beBad(func(tb testing.TB) { be.Nonzero(tb, time.Time{}.Local()) })
	beBad(func(tb testing.TB) { be.Zero(tb, []string{""}) })
	beBad(func(tb testing.TB) { be.Nonzero(tb, []string(nil)) })
	beBad(func(tb testing.TB) { be.NilErr(tb, errors.New("")) })
	beBad(func(tb testing.TB) { be.True(tb, false) })
	beBad(func(tb testing.TB) { be.False(tb, true) })
	beBad(func(tb testing.TB) { be.EqualLength(tb, 1, ch) })
	ch <- 1
	beBad(func(tb testing.TB) { be.EqualLength(tb, 0, ch) })
	close(ch)
	beBad(func(tb testing.TB) { be.EqualLength(tb, 1, ch) })
}
