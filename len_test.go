package be_test

import (
	"testing"

	"github.com/carlmjohnson/be"
)

func TestLen(t *testing.T) {
	// Make sure integers aren't treated as rangeable
	be.Nonzero(t, be.Panicked(func() {
		be.EqualLength(t, 0, 0)
	}))
}
