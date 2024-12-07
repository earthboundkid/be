package be_test

import "github.com/carlmjohnson/be"

func ExamplePanicked() {
	// mock *testing.T for example purposes
	t := &mockingT{}

	divide := func(num, denom int) int {
		return num / denom
	}

	// Test that division by zero panics
	be.Nonzero(t, be.Panicked(func() {
		divide(1, 0)
	}))

	// Because a panic fails a test by default,
	// testing that an operation does not panic is less necessary,
	// but may be helpful in a table test.
	denom := 2
	panicVal := be.Panicked(func() {
		divide(1, denom)
	})
	wantPanic := denom == 0
	be.Equal(t, wantPanic, panicVal != nil)

	// Output:
}
