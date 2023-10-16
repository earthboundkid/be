// Package be is a minimalist test assertion helper library.
//
// # Philosophy
//
// Tests usually should not fail. When they do fail, the failure should be repeatable. Therefore, it doesn't make sense to spend a lot of time writing good test messages. (This is unlike error messages, which should happen fairly often, and in production, irrepeatably.) Package be is designed to simply fail a test quickly and quietly if a condition is not met with a reference to the line number of the failing test. If the reason for having the test is not immediately clear from context, you can write a comment, like normal code. If you do need more extensive reporting to figure out why a test is failing, use [be.DebugLog] or [be.Debug] to capture more information.
//
// Most tests just need simple equality testing, which is handled by [be.Equal] (for comparable types), [be.AllEqual] (for slices of comparable types), and [be.DeepEqual] (which relies on [reflect.DeepEqual]). Another common test is that a string or byte slice should contain or not some substring, which is handled by [be.In] and [be.NotIn]. Rather than package be providing every possible test helper, you are encouraged to write your own advanced helpers for use with [be.True], while package be takes away the drudgery of writing yet another simple func nilErr(t *testing.T, err) { ... }.
//
// The [github.com/carlmjohnson/be/testfile] subpackage has functions that make it easy to write file-based tests that ensure that the output of some transformation matches a golden file. Subtests can automatically be run for all files matching a glob pattern, such as testfile.Run(t, "testdata/*/input.txt", ...). If the test fails, the failure output will be written to a file, such as "testdata/basic-test/-failed-output.txt", and then the output can be examined via diff testing with standard tools. Set the environmental variable TESTFILE_UPDATE to update the golden file.
//
// Every test in the be package requires a [testing.TB] as its first argument. There are various clever ways to get the testing.TB implicitly,[*] but package be is designed to be simple and explicit, so it's easiest to just always pass in a testing.TB the boring way.
//
// [*]: https://dave.cheney.net/2019/12/08/dynamically-scoped-variables-in-go
package be
