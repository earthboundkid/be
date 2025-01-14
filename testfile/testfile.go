// Package testfile has test helpers that work by comparing files.
package testfile

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Ext return path with its current extension stripped and ext added.
// It treats ext with and without a leading dot the same for simplicity of operation.
// If ext is "", path is returned with its current extension stripped off.
func Ext(path, ext string) string {
	currExt := filepath.Ext(path)
	path = strings.TrimSuffix(path, currExt)
	if !strings.HasPrefix(ext, ".") && ext != "" {
		ext = "." + ext
	}
	return path + ext
}

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
//
// If the environmental variable TESTFILE_UPDATE is set,
// an unequal file will be overwritten with gotStr,
// but the test will still fail
// to prevent accidental updates from going unnoticed.
func Equal(t testing.TB, wantFile, gotStr string) {
	t.Helper()
	equal(t, wantFile, gotStr, false)
}

// Equalish is like Equal,
// but it uses strings.TrimSpace before testing for equality.
//
// If the environmental variable TESTFILE_UPDATE is set,
// an unequal file will be overwritten with gotStr,
// but the test will still fail
// to prevent accidental updates from going unnoticed.
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
	dir, name := filepath.Split(wantFile)
	badFile := filepath.Join(dir, "-failed-"+name)
	if w == gotStr {
		// Remove lingering test failures
		_ = os.Remove(badFile)
		return
	}
	if os.Getenv("TESTFILE_UPDATE") != "" {
		// Remove lingering test failures
		_ = os.Remove(badFile)
		Write(t, wantFile, gotStr)
		t.Errorf("updated contents of %s", wantFile)
		return
	}
	Write(t, badFile, gotStr)
	t.Fatalf("contents of %s != %s", wantFile, badFile)
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
// The contents of wantFile must have two spaces for indentation.
// EqualJSON just uses string equality
// and does not test for JSON equivalency.
// If they are not equal, it writes out a file with the contents of v and calls t.Fatalf.
// If there is an error, it calls t.Fatalf.
//
// If the environmental variable TESTFILE_UPDATE is set,
// an unequal file will be overwritten with v,
// but the test will still fail
// to prevent accidental updates from going unnoticed.
func EqualJSON(t testing.TB, wantFile string, v any) {
	t.Helper()
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(v)
	if err != nil {
		t.Fatalf("marshaling: %v", err)
		return
	}
	Equalish(t, wantFile, buf.String())
}

// WriteJSON writes v as JSON to a file at path.
// The JSON is created using two spaces for indentation.
// If there is an error, it calls t.Fatalf.
func WriteJSON(t testing.TB, path string, v any) {
	t.Helper()
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(v)
	if err != nil {
		t.Fatalf("marshaling: %v", err)
		return
	}
	Write(t, path, buf.String())
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
