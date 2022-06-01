# Be [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/be.svg)](https://pkg.go.dev/github.com/carlmjohnson/be) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/be)](https://goreportcard.com/report/github.com/carlmjohnson/be) [![Gocover.io](https://gocover.io/_badge/github.com/carlmjohnson/be)](https://gocover.io/github.com/carlmjohnson/be) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
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
Tests usually should not fail. When they do fail, the failure should be repeatable. Therefore, it doesn't make sense to spend a lot of time writing good test messages. (This is unlike error messages, which should happen fairly often, and in production, irrepeatably.) Package be is designed to simply fail a test quickly and quietly if a condition is not met with a reference to the line number of the failing test. If you do need more extensive reporting to figure out why a test is failing, use `be.DebugLog` or `be.Debug` to capture more information, for example by saving the failing data to disk for diffing and examination.

Most tests just need simple equality testing, which is handled by `be.Equal` (for comparable types), `be.AllEqual` (for slices of comparable types), and `be.DeepEqual` (which relies on `reflect.DeepEqual`). Another common test is that a string or byte slice should contain or not some substring, which is handled by `be.In` and `be.NotIn`. Rather than package be providing every possible test helper, you are encouraged to write your own advanced helpers for use with `be.True`, while package be takes away the drudgery of writing yet another simple `func nilErr(t *testing.T, err) { ... }`.

Every test in the be package requires a `testing.TB` as its first argument. There are various [clever ways to get the testing.TB implicitly](https://dave.cheney.net/2019/12/08/dynamically-scoped-variables-in-go), but package be is designed to be simple and explicit, so it's easiest to just always pass in a testing.TB the boring way.
