// Package testfile has test helpers that work by comparing files.
package testfile

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Read returns the contents of file.
// It calls t.Fatalf if there is an error.
func Read(t testing.TB, name string) string {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return string(b)
}

// Write the data to a file with 0644 permission bits.
// It attempts to create directories as needed.
// It calls t.Fatalf if there is an error.
func Write(t testing.TB, name, data string) {
	t.Helper()
	dir := filepath.Dir(name)
	_ = os.MkdirAll(dir, 0700)
	err := os.WriteFile(name, []byte(data), 0644)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

// Equal tests whether gotStr is equal to the contents of wantFile.
// If they are not equal,
// it writes gotStr to a file with -bad added to its name
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
	ext := filepath.Ext(wantFile)
	base := strings.TrimSuffix(wantFile, ext)
	name := base + "-bad" + ext
	Write(t, name, gotStr)
	t.Fatalf("contents of %s != %s", wantFile, name)
}

// ReadJSON attempts to unmarshal the contents of a file into v.
// If there is an error, it calls t.Fatalf.
func ReadJSON(t testing.TB, name string, v any) {
	t.Helper()
	s := Read(t, name)
	if err := json.Unmarshal([]byte(s), v); err != nil {
		t.Fatalf("unmarshal %s: %v", name, err)
	}
}

// EqualJSON tests whether v is equal to the contents of wantFile when mashaled as JSON.
// The JSON must be created with json.MarshalIndent and have two spaces as a prefix.
// If they are not equal, it writes a file with the contents of v and calls t.Fatalf.
// If there is an error, it calls t.Fatalf.
func EqualJSON(t testing.TB, wantFile string, v any) {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("marshaling: %v", err)
		return
	}
	Equalish(t, wantFile, string(b))
}

// WriteJSON writes v as to name as JSON
// with json.MarshalIndent and has two spaces as a prefix.
// If there is an error, it calls t.Fatalf.
func WriteJSON(t testing.TB, name string, v any) {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("marshaling: %v", err)
		return
	}
	Write(t, name, string(b))
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
