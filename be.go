// Package be is a minimalist test assertion helper library.
package be

import (
	"reflect"
	"testing"
)

// Eq calls t.Fatal if want != got.
func Eq[T comparable](t testing.TB, want, got T) {
	t.Helper()
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

// NotEq calls t.Fatal if got == bad.
func NotEq[T comparable](t testing.TB, bad, got T) {
	t.Helper()
	if got == bad {
		t.Fatalf("got: %v", got)
	}
}

// EqSlice calls t.Fatal if want != got.
func EqSlice[T comparable](t testing.TB, want, got []T) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("len(want): %d; len(got): %v", len(want), len(got))
		return
	}
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("want: %v; got: %v", want, got)
			return
		}
	}
}

// Zero calls t.Fatal if value != the zero value for T.
func Zero[T any](t testing.TB, value T) {
	t.Helper()
	if truthy(value) {
		t.Fatalf("got: %v", value)
	}
}

// Zero calls t.Fatal if value == the zero value for T.
func NonZero[T any](t testing.TB, value T) {
	t.Helper()
	if !truthy(value) {
		t.Fatalf("got: %v", value)
	}
}

func truthy[T any](v T) bool {
	switch m := any(v).(type) {
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&v)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}
