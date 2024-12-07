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
	for _, denom := range []int{-1, 0, 1, 1_000} {
		shouldPanic := denom == 0
		panicVal := be.Panicked(func() {
			divide(1, denom)
		})
		be.Equal(t, shouldPanic, panicVal != nil)
	}

	// Output:
}
