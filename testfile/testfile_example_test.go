package testfile_test

import (
	"fmt"
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
	testfile.Write(t, "example/upper.txt", got)
	// Read files
	file := testfile.Read(t, "example/upper.txt")
	// Do some testing
	be.Equal(t, file, got)
	// Use the equality helper
	testfile.Equal(t, "example/upper.txt", got)
	// If the files aren't equal,
	// got will be written to example/-failed-upper.txt

	// Or use JSON helpers for complex structs
	type testcase struct {
		Input, Output string
	}
	s := testcase{input, got}
	// Write out pretty printed JSON
	testfile.WriteJSON(t, "example/upper.json", s)
	// Read from a JSON file
	var s2 testcase
	testfile.ReadJSON(t, "example/upper.json", &s2)
	be.Equal(t, s, s2)
	// Test that s equals the contents of a file when serialized
	testfile.EqualJSON(t, "example/upper.json", s)

	// Output:
}

func ExampleRun() {
	_ = func(t *testing.T) {
		// For each .txt file, start a sub-test
		testfile.Run(t, "testdata/*.txt", func(t *testing.T, path string) {
			// Read the file
			input := testfile.Read(t, path)

			// Do some conversion on it
			type myStruct struct{ Whatever string }
			got := myStruct{strings.ToUpper(input)}

			// See if the struct is equivalent to a .json file
			wantFile := testfile.Ext(path, ".json")
			testfile.EqualJSON(t, wantFile, got)

			// If it's not equivalent,
			// the got struct will be dumped
			// to a file named testdata/-failed-test-name.json
		})
	}
	// Output:
}

func ExampleExt() {
	for _, path := range []string{
		"foo.txt",
		"foo.bar/spam",
	} {
		fmt.Println(testfile.Ext(path, ""))
		fmt.Println(testfile.Ext(path, ".json"))
		fmt.Println(testfile.Ext(path, "gob"))
	}
	// Output:
	// foo
	// foo.json
	// foo.gob
	// foo.bar/spam
	// foo.bar/spam.json
	// foo.bar/spam.gob
}
