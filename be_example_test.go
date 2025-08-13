package be_test

import (
	"errors"
	"strings"

	"github.com/carlmjohnson/be"
)

func Example() {
	// mock *testing.T for example purposes
	t := be.Relaxed(&mockingT{})

	be.Equal(t, "hello", "world")     // bad
	be.Equal(t, "goodbye", "goodbye") // good

	be.Unequal(t, "hello", "world")     // good
	be.Unequal(t, "goodbye", "goodbye") // bad

	s := []int{1, 2, 3}
	be.AllEqual(t, []int{1, 2, 3}, s) // good
	be.AllEqual(t, []int{3, 2, 1}, s) // bad

	var err error
	be.NilErr(t, err)  // good
	be.Nonzero(t, err) // bad
	err = errors.New("(O_o)")
	be.NilErr(t, err)  // bad
	be.Nonzero(t, err) // good

	type mytype string
	var mystring mytype = "hello, world"
	be.In(t, "world", mystring)                 // good
	be.In(t, "World", mystring)                 // bad
	be.NotIn(t, "\x01", []byte("\a\b\x00\r\t")) // good
	be.NotIn(t, "\x00", []byte("\a\b\x00\r\t")) // bad

	seq := strings.FieldsSeq("1 2 3 4")
	be.EqualLength(t, 4, seq)     // good
	be.EqualLength(t, 1, seq)     // bad
	be.AtLeastLength(t, 1, seq)   // good
	be.AtLeastLength(t, 5, seq)   // bad
	be.AtLeastLength(t, 3, "123") // good
	be.AtLeastLength(t, 4, "123") // bad

	// Output:
	// want: hello; got: world
	// got: goodbye
	// want: [3 2 1]; got: [1 2 3]
	// got: <nil>
	// got: (O_o)
	// "World" not in "hello, world"
	// "\x00" in "\a\b\x00\r\t"
	// want len(seq) == 1; got at least 2
	// want len(seq) >= 5; got 4
	// want len(seq) >= 4; got 3
}
