package be_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/carlmjohnson/be"
)

type mockingT struct{ *testing.T }

func (_ mockingT) Helper() {}

func (_ mockingT) Fatalf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func Example() {
	// mock *testing.T for example purposes
	var t mockingT

	be.Eq(t, "hello", "world")     // bad
	be.Eq(t, "goodbye", "goodbye") // good

	be.NotEq(t, "hello", "world")     // good
	be.NotEq(t, "goodbye", "goodbye") // bad

	s := []int{1, 2, 3}
	be.EqSlice(t, []int{1, 2, 3}, s) // good
	be.EqSlice(t, []int{3, 2, 1}, s) // bad

	var err error
	be.Zero(t, err)    // good
	be.NonZero(t, err) // bad
	err = errors.New("(O_o)")
	be.Zero(t, err)    // bad
	be.NonZero(t, err) // good

	// Output:
	// want: hello; got: world
	// got: goodbye
	// want: [3 2 1]; got: [1 2 3]
	// got: <nil>
	// got: (O_o)
}
