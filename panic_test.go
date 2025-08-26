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

func TestMatch(t *testing.T) {
	// Make sure bad regexp patterns panic
	pval := be.Panicked(func() {
		be.Match(t, `\`, "")
	})
	be.Nonzero(t, pval)
	s, ok := pval.(string)
	be.True(t, ok)
	be.Match(t, `^regexp: Compile\(`, s)
}
