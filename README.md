# Be [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/be.svg)](https://pkg.go.dev/github.com/carlmjohnson/be) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/be)](https://goreportcard.com/report/github.com/carlmjohnson/be) [![Coverage Status](https://coveralls.io/repos/github/carlmjohnson/be/badge.svg)](https://coveralls.io/github/carlmjohnson/be) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
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
be.NotIn(t, "\x01", []byte("\a\b\x00\r\t")) // good
be.NotIn(t, "\x00", []byte("\a\b\x00\r\t")) // bad
// t.Fatal("\x00" in "\a\b\x00\r\t")
```

## Philosophy
Tests usually should not fail. When they do fail, the failure should be repeatable. Therefore, it doesn't make sense to spend a lot of time writing good test messages. (This is unlike error messages, which should happen fairly often, and in production, irrepeatably.) Package be is designed to simply fail a test quickly and quietly if a condition is not met with a reference to the line number of the failing test. If the reason for having the test is not immediately clear from context, you can write a comment, like normal code. If you do need more extensive reporting to figure out why a test is failing, use `be.DebugLog` or `be.Debug` to capture more information.

Most tests just need simple equality testing, which is handled by `be.Equal` (for comparable types), `be.AllEqual` (for slices of comparable types), and `be.DeepEqual` (which relies on `reflect.DeepEqual`). Another common test is that a string or byte slice should contain or not some substring, which is handled by `be.In` and `be.NotIn`. Rather than package be providing every possible test helper, you are encouraged to write your own advanced helpers for use with `be.True`, while package be takes away the drudgery of writing yet another simple `func nilErr(t *testing.T, err) { ... }`.

The testfile subpackage has functions that make it easy to write file-based tests that ensure that the output of some transformation matches a [golden file](https://softwareengineering.stackexchange.com/questions/358786/what-are-golden-files). Subtests can automatically be run for all files matching a glob pattern, such as `testfile.Run(t, "testdata/*/input.txt", ...)`. If the test fails, the failure output will be written to a file, such as "testdata/basic-test/-failed-output.txt", and then the output can be examined via diff testing with standard tools. Writing failed outputs also removes the need for an "update" mode that some other golden testfile packages have, since you can just run the test once, examine the failed output, and renamed the file if it is actually correct.

Every tool in the be module requires a `testing.TB` as its first argument. There are various [clever ways to get the testing.TB implicitly](https://dave.cheney.net/2019/12/08/dynamically-scoped-variables-in-go), but package be is designed to be simple and explicit, so it's easiest to just always pass in a testing.TB the boring way.
