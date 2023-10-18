package testfile_test

import (
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/carlmjohnson/be/testfile"
)

func runPaths(t *testing.T, inpath string) []string {
	var paths []string
	testfile.Run(t, inpath, func(t *testing.T, path string) {
		paths = append(paths, path)
	})
	return paths
}

func TestRun(t *testing.T) {
	cases := []struct {
		InPath    string
		WantFound string
	}{
		{"testdata/*.glorp", ""},
		{"testdata/*.json", "testdata/example.json"},
	}
	for _, tc := range cases {
		got := runPaths(t, tc.InPath)
		be.Equal(t, tc.WantFound, strings.Join(got, ","))
	}
}

func TestEqualJSON(t *testing.T) {
	testfile.EqualJSON(t, "testdata/example.json", struct {
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

func TestSetEnv(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "example.txt")
	failedPath := filepath.Join(dir, "-failed-example.txt")

	// If testfile.Equal fails
	testfile.Write(t, path, "1")
	be.False(t, runTest(func(t *testing.T) {
		testfile.Equal(t, path, "2")
	}))
	// it writes a -failed file.
	testfile.Equal(t, failedPath, "2")
	// If testfile.Equal succeeds,
	testfile.Equal(t, path, "1")
	// it erases the -failed file.
	testfile.Equal(t, failedPath, "")
	// If TESTFILE_UPDATE is set,
	t.Setenv("TESTFILE_UPDATE", "ON")
	// the test still fails
	be.False(t, runTest(func(t *testing.T) {
		testfile.Equal(t, path, "3")
	}))
	// but it doesn't write a failed path,
	testfile.Equal(t, failedPath, "")
	// and does update the file.
	testfile.Equal(t, path, "3")
}
