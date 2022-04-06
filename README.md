# Be [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/be.svg)](https://pkg.go.dev/github.com/carlmjohnson/be)
Package be is the minimalist testing helper for Go 1.18+.

Inspired by [Mat Ryer](https://github.com/matryer/is) and [Alex Edwards](https://www.alexedwards.net/blog/easy-test-assertions-with-go-generics).

```go
be.Equal(t, "hello", "world")     // bad
// t.Fatal("want: hello; got: world")
be.Equal(t, "goodbye", "goodbye") // good

be.Unequal(t, "hello", "world")     // good
be.Unequal(t, "goodbye", "goodbye") // bad
// t.Fatal("got: goodbye")

s := []int{1, 2, 3}
be.AllEqual(t, []int{1, 2, 3}, s) // good
be.AllEqual(t, []int{3, 2, 1}, s) // bad
// t.Fatal("want: [3 2 1]; got: [1 2 3]")

var err error
be.NilErr(t, err)   // good
be.Nonzero(t, err) // bad
// t.Fatal("got: <nil>")
err = errors.New("(O_o)")
be.NilErr(t, err)   // bad
// t.Fatal("got: (O_o)")
be.Nonzero(t, err) // good

be.In(t, "world", "hello, world") // good
be.In(t, "World", "hello, world") // bad
// t.Fatal("World" not in "hello, world")
be.In(t, "\x00", []byte("\a\b\x00\r\t")) // good
be.In(t, "\x01", []byte("\a\b\x00\r\t")) // bad
// t.Fatal("\x01" not in "\a\b\x00\r\t")

```
