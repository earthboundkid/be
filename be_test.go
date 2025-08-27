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
	okayTests := []func(tb testing.TB){
		func(tb testing.TB) { be.Zero(tb, time.Time{}.Local()) },
		func(tb testing.TB) { be.Zero(tb, []string(nil)) },
		func(tb testing.TB) { be.Nonzero(tb, []string{""}) },
		func(tb testing.TB) { be.NilErr(tb, nil) },
		func(tb testing.TB) { be.True(tb, true) },
		func(tb testing.TB) { be.False(tb, false) },
		func(tb testing.TB) { be.EqualLength(tb, 0, map[int]int{}) },
		func(tb testing.TB) { be.EqualLength(tb, 1, map[int]int{1: 1}) },
		func(tb testing.TB) {
			ch := make(chan int, 1)
			be.EqualLength(tb, 0, ch)
		},
		func(tb testing.TB) {
			ch := make(chan int, 1)
			ch <- 1
			be.EqualLength(tb, 1, ch)
		},
		func(tb testing.TB) {
			seq2 := maps.All(map[int]int{1: 1})
			be.EqualLength(tb, 1, seq2)
		},
		func(tb testing.TB) {
			be.In(tb, "world", "Hello, world!")
			be.NotIn(tb, "\x01", []byte("\a\b\x00\r\t"))
		},
	}

	for _, test := range okayTests {
		var buf strings.Builder
		tb := &testingTB{w: &buf}
		test(tb)
		if tb.failed {
			t.Fatal("failed too soon")
		}
		if buf.String() != "" {
			t.Fatal("wrote too much")
		}
	}

	badTests := []func(tb testing.TB){
		func(tb testing.TB) { be.AllEqual(tb, []string{}, []string{""}) },
		func(tb testing.TB) { be.Nonzero(tb, time.Time{}.Local()) },
		func(tb testing.TB) { be.Zero(tb, []string{""}) },
		func(tb testing.TB) { be.Nonzero(tb, []string(nil)) },
		func(tb testing.TB) { be.NilErr(tb, errors.New("")) },
		func(tb testing.TB) { be.True(tb, false) },
		func(tb testing.TB) { be.False(tb, true) },
		func(tb testing.TB) {
			seq2 := maps.All(map[int]int{1: 1})
			be.EqualLength(tb, 0, seq2)
		},
		func(tb testing.TB) {
			ch := make(chan int, 1)
			be.EqualLength(tb, 1, ch)
		},
		func(tb testing.TB) {
			ch := make(chan int, 1)
			ch <- 1
			be.EqualLength(tb, 0, ch)
		},
		func(tb testing.TB) {
			ch := make(chan int, 1)
			close(ch)
			be.EqualLength(tb, 1, ch)
		},
		func(tb testing.TB) {
			be.In(tb, "World", "Hello, world!")
		},
		func(tb testing.TB) {
			be.NotIn(tb, "\x00", []byte("\a\b\x00\r\t"))
		},
	}

	for _, test := range badTests {
		var buf strings.Builder
		tb := &testingTB{w: &buf}
		test(tb)
		if !tb.failed {
			t.Fatal("did not fail")
		}
		if buf.String() == "" {
			t.Fatal("wrote too little")
		}
	}
}
