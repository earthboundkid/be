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
	exampleRun := func(t *testing.T) {
		// For each .txt file, start a sub-test
		testfile.Run(t, "example-run/*.txt", func(t *testing.T, path string) {
			type testdata struct{ Input, Output string }
			// Read the file
			input := testfile.Read(t, path)
			// Do some conversion on it
			output := strings.ToUpper(input)
			got := testdata{input, output}

			// See if the struct is equivalent to a .json file
			wantFile := strings.TrimSuffix(path, ".txt") + ".json"
			testfile.EqualJSON(t, wantFile, got)
			fmt.Printf("Test of %s passed!\n", path)
		})
	}

	all := func(pat, str string) (bool, error) { return true, nil }
	testing.RunTests(all, []testing.InternalTest{{"testfile.Run", exampleRun}})

	// Output:
	// Test of example-run/hello.txt passed!
}
