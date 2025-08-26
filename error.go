package be

import (
	"errors"
	"testing"
)

// ErrorIs calls t.Fatalf if got is not want according to [errors.Is].
func ErrorIs(t testing.TB, want, got error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Fatalf("got errors.Is(%v, %v) == false", got, want)
	}
}

// ErrorIs calls t.Fatalf if got cannot be assigned to want by [errors.As].
func ErrorAs[T error](t testing.TB, want *T, got error) {
	t.Helper()
	if !errors.As(got, want) {
		t.Fatalf("got errors.As(%v, %T) == false", got, want)
	}
}
