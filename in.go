package be

import (
	"regexp"
	"testing"
)

// In calls t.Fatalf if needle is not contained in the string or []byte haystack.
//
// Deprecated: Use Match(t, regexp.QuoteMeta(needle), haystack).
//
//go:fix inline
func In[byteseq ~string | ~[]byte](t testing.TB, needle string, haystack byteseq) {
	Match(t, regexp.QuoteMeta(needle), haystack)
}

// NotIn calls t.Fatalf if needle is contained in the string or []byte haystack.
//
// Deprecated: Use NoMatch(t, regexp.QuoteMeta(needle), haystack).
//
//go:fix inline
func NotIn[byteseq ~string | ~[]byte](t testing.TB, needle string, haystack byteseq) {
	NoMatch(t, regexp.QuoteMeta(needle), haystack)
}
