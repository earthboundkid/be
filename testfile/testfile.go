// Package testfile has test helpers that work by comparing files.
package testfile

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Read returns the contents of file at path.
// It calls t.Fatalf if there is an error.
func Read(t testing.TB, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return string(b)
}

// Write the data to a file at path with 0644 permission bits.
// It attempts to create directories as needed.
// It calls t.Fatalf if there is an error.
func Write(t testing.TB, path, data string) {
	t.Helper()
	dir := filepath.Dir(path)
	_ = os.MkdirAll(dir, 0700)
	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

// Equal tests whether gotStr is equal to the contents of wantFile.
// If they are not equal,
// it writes gotStr to a file with -failed- prepended to its name
// and calls t.Fatalf.
func Equal(t testing.TB, wantFile, gotStr string) {
	t.Helper()
	equal(t, wantFile, gotStr, false)
}

// Equalish is like Equal,
// but it uses strings.TrimSpace before testing for equality.
func Equalish(t testing.TB, wantFile, gotStr string) {
	t.Helper()
	equal(t, wantFile, gotStr, true)
}

func equal(t testing.TB, wantFile, gotStr string, trim bool) {
	t.Helper()
	b, err := os.ReadFile(wantFile)
	switch {
	case err == nil, os.IsNotExist(err):
	case err != nil:
		t.Fatalf("%v", err)
		return
	}
	w := string(b)
	if trim {
		w = strings.TrimSpace(w)
		gotStr = strings.TrimSpace(gotStr)
	}
	if w == gotStr {
		return
	}
	dir, name := filepath.Split(wantFile)
	name = filepath.Join(dir, "-failed-"+name)
	Write(t, name, gotStr)
	t.Fatalf("contents of %s != %s", wantFile, name)
}

// ReadJSON attempts to unmarshal the contents of a file at path into v.
// If there is an error, it calls t.Fatalf.
func ReadJSON(t testing.TB, path string, v any) {
	t.Helper()
	s := Read(t, path)
	if err := json.Unmarshal([]byte(s), v); err != nil {
		t.Fatalf("unmarshal %s: %v", path, err)
	}
}

// EqualJSON tests whether
// when v is mashaled as JSON,
// it is equal to the contents of wantFile.
// EqualJSON ignores whitespace in wantFile,
// but is sensitive to the order of keys.
// If they are not equal, it writes out a file with the contents of v and calls t.Fatalf.
// If there is an error, it calls t.Fatalf.
func EqualJSON(t testing.TB, wantFile string, v any) {
	t.Helper()
	got, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("marshaling v: %v", err)
		return
	}
	b, err := os.ReadFile(wantFile)
	if err != nil {
		t.Fatalf("reading wantFile: %v", err)
		return
	}
	var buf bytes.Buffer
	buf.Grow(len(b))
	if err = json.Indent(&buf, b, "", "  "); err != nil {
		t.Fatalf("indenting wantFile: %v", err)
		return
	}
	gotStr := string(bytes.TrimSpace(got))
	want := strings.TrimSpace(buf.String())
	if gotStr == want {
		return
	}
	dir, name := filepath.Split(wantFile)
	name = filepath.Join(dir, "-failed-"+name)
	Write(t, name, gotStr)
	t.Fatalf("contents of %s != %s", wantFile, name)
}

// WriteJSON writes v as JSON to a file at path.
// The JSON is created with json.MarshalIndent using two spaces for indentation.
// If there is an error, it calls t.Fatalf.
func WriteJSON(t testing.TB, path string, v any) {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("marshaling: %v", err)
		return
	}
	Write(t, path, string(b))
}

// Run a subtest for each file matching the provided glob pattern.
func Run(t *testing.T, glob string, f func(t *testing.T, match string)) {
	t.Helper()
	matches, err := filepath.Glob(glob)
	if err != nil {
		t.Fatalf("%v", err)
		return
	}
	for i := range matches {
		match := matches[i]
		name := filepath.Base(match)
		t.Run(name, func(t *testing.T) {
			f(t, match)
		})
	}
}
