package be

import (
	"reflect"
	"regexp"
	"testing"
)

// Match calls t.Fatalf if got does not match the [regexp] pattern.
//
// The pattern must compile.
func Match[byteseq ~string | ~[]byte](t testing.TB, pattern string, got byteseq) {
	t.Helper()
	reg := regexp.MustCompile(pattern)
	if !match(reg, got) {
		t.Fatalf("/%s/ !~ %q", pattern, got)
	}
}

func match[byteseq ~string | ~[]byte](reg *regexp.Regexp, got byteseq) bool {
	switch rv := reflect.ValueOf(got); rv.Kind() {
	case reflect.String:
		return reg.MatchString(rv.String())
	case reflect.Slice:
		return reg.Match(rv.Bytes())
	}
	panic("unreachable")
}
