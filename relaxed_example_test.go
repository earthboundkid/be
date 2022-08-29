package be_test

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func ExampleRelaxed() {
	// mock *testing.T for example purposes
	t := &mockingT{}

	t.Run("dies on first error", func(*testing.T) {
		be.Equal(t, 1, 2)
		be.Equal(t, 3, 4)
	})

	t.Run("shows multiple errors", func(*testing.T) {
		relaxedT := be.Relaxed(t)
		be.Equal(relaxedT, 5, 6)
		be.Equal(relaxedT, 7, 8)
	})
	// Output:
	// want: 1; got: 2
	// want: 5; got: 6
	// want: 7; got: 8
}
