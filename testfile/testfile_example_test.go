package testfile_test

import (
	"strings"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/carlmjohnson/be/testfile"
)

func Example() {
	t := &testing.T{}

	// Make some test data
	input := "Hello, World!"
	got := strings.ToUpper(input)
	// Write files
	testfile.Write(t, "upper.txt", got)
	// Read files
	file := testfile.Read(t, "upper.txt")
	// Do some testing
	be.Equal(t, file, got)
	// Use the equality helper
	testfile.Equal(t, "upper.txt", got)

	// Or use JSON helpers for complex structs
	type testcase struct {
		Input, Output string
	}
	s := testcase{input, got}
	// Write out pretty printed JSON
	testfile.WriteJSON(t, "upper.json", s)
	// Read from a JSON file
	var s2 testcase
	testfile.ReadJSON(t, "upper.json", &s2)
	be.Equal(t, s, s2)
	// Test that s equals the contents of a file when serialized
	testfile.EqualJSON(t, "upper.json", s)

	// Output:
}
