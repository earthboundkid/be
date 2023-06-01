package testfile_test

import (
	"math"
	"path/filepath"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/carlmjohnson/be/testfile"
)

func TestRunEqualJSON(t *testing.T) {
	testfile.Run(t, "testdata/*.json", func(t *testing.T, path string) {
		testfile.EqualJSON(t, path, struct {
			Data any `json:"data"`
		}{
			Data: []struct {
				Field string `json:"field"`
				Value int    `json:"value"`
			}{
				{"foo", 1},
				{"bar", 2},
			},
		})
	})
}

func runTest(test func(*testing.T)) bool {
	return testing.RunTests(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{{"test", test}})
}

func TestCases(t *testing.T) {
	dir := t.TempDir()
	join := func(s string) string {
		return filepath.Join(dir, s)
	}
	cases := []struct {
		name string
		want bool
		f    func(*testing.T)
	}{
		{"bad Read", false, func(t *testing.T) {
			_ = testfile.Read(t, join("nope.txt"))
		}},
		{"bad Write", false, func(t *testing.T) {
			testfile.Write(t, join("\x00"), "")
		}},
		{"bad Equal", false, func(t *testing.T) {
			f := join("bad equal")
			testfile.Write(t, f, "bad")
			testfile.Equal(t, f, "not equal")
		}},
		{"bad Equal file", false, func(t *testing.T) {
			testfile.Equal(t, "/", "not equal")
		}},
		{"bad ReadJSON", false, func(t *testing.T) {
			f := join("bad-read-json.txt")
			testfile.Write(t, f, "xxx")
			var v any
			testfile.ReadJSON(t, f, v)
		}},
		{"bad WriteJSON", false, func(t *testing.T) {
			testfile.WriteJSON(t, join("out.txt"), math.NaN())
		}},
		{"bad EqualJSON", false, func(t *testing.T) {
			testfile.EqualJSON(t, join("out.txt"), math.NaN())
		}},
		{"good WriteJSON", true, func(t *testing.T) {
			testfile.WriteJSON(t, join("out.txt"), 1.0)
		}},
		{"bad Run glob", false, func(t *testing.T) {
			testfile.Run(t, "[]", func(t *testing.T, match string) {})
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			be.Equal(t, tc.want, runTest(tc.f))
		})
	}
}
