package be_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/carlmjohnson/be"
)

type mockingT struct{ *testing.T }

func (_ mockingT) Helper() {}

func (_ mockingT) Fatal(args ...any) {
	fmt.Println(args...)
}

func Example() {
	// mock *testing.T for example purposes
	var t mockingT

	be.Equal(t, "hello", "world")     // bad
	be.Equal(t, "goodbye", "goodbye") // good

	be.Unequal(t, "hello", "world")     // good
	be.Unequal(t, "goodbye", "goodbye") // bad

	s := []int{1, 2, 3}
	be.AllEqual(t, []int{1, 2, 3}, s) // good
	be.AllEqual(t, []int{3, 2, 1}, s) // bad

	var err error
	be.Zero(t, err)    // good
	be.Nonzero(t, err) // bad
	err = errors.New("(O_o)")
	be.Zero(t, err)    // bad
	be.Nonzero(t, err) // good

	// Output:
	// be_example_test.go:23 want: hello; got: world
	// be_example_test.go:27 got: goodbye
	// be_example_test.go:31 want: [3 2 1]; got: [1 2 3]
	// be_example_test.go:35 got: <nil>
	// be_example_test.go:37 got: (O_o)
}
