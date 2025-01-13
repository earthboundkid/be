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
	for _, testcase := range []struct {
		num, denom, want int
		shouldPanic      bool
	}{
		{0, 1, 0, false},
		{1, 1, 1, false},
		{1, 0, 0xbadc0ffee, true},
		{0, 0, 0xbadc0ffee, true},
	} {
		got := 0xbadc0ffee
		panicVal := be.Panicked(func() {
			got = divide(testcase.num, testcase.denom)
		})
		be.Equal(t, testcase.want, got)
		be.Equal(t, testcase.shouldPanic, panicVal != nil)
	}
	// Output:
}
