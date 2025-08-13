package be

import (
	"iter"
	"reflect"
	"slices"
	"testing"
)

// EqualLength calls t.Fatalf if seq has a length that is not exactly want.
// The type of seq must be array, array pointer, slice, map, string, channel, [iter.Seq], or [iter.Seq2].
// For channels and iterators, the values are consumed to get the sequence length.
func EqualLength(t testing.TB, want int, seq any) {
	t.Helper()
	getLen(t, seq, want, false)
}

func getLen(t testing.TB, seq any, n int, atLeast bool) {
	rv := reflect.ValueOf(seq)
	rt := rv.Type()
	kind := rt.Kind()
	switch {
	case slices.Contains([]reflect.Kind{
		reflect.Array, reflect.Slice, reflect.Map, reflect.String,
	}, kind) ||
		(kind == reflect.Pointer && rt.Elem().Kind() == reflect.Array):
		compareN(t, n, rv.Len(), atLeast)
	case kind == reflect.Chan:
		eqchan(t, rv, n, atLeast)
	case kind == reflect.Func && rt.CanSeq():
		eqseq(t, rv, n, atLeast)
	case kind == reflect.Func && rt.CanSeq2():
		eqseq2(t, rv, n, atLeast)
	default:
		panic("seq must be a non-integer rangeable type")
	}
}

func eqchan(t testing.TB, rv reflect.Value, n int, atLeast bool) {
	i := 0
	for ; i <= n; i++ {
		chosen, _, recvOK := reflect.Select([]reflect.SelectCase{
			{Dir: reflect.SelectRecv, Chan: rv},
			{Dir: reflect.SelectDefault},
		})

		if chosen == 1 || !recvOK { // default case or channel closed
			break
		}
	}
	compareN(t, n, i, atLeast)
}

func eqseq(t testing.TB, rv reflect.Value, n int, atLeast bool) {
	next, stop := iter.Pull(rv.Seq())
	defer stop()
	i := 0
	for ; i <= n; i++ {
		if _, ok := next(); !ok {
			break
		}
	}
	compareN(t, n, i, atLeast)
}

func eqseq2(t testing.TB, rv reflect.Value, n int, atLeast bool) {
	next, stop := iter.Pull2(rv.Seq2())
	defer stop()
	i := 0
	for ; i <= n; i++ {
		if _, _, ok := next(); !ok {
			break
		}
	}
	compareN(t, n, i, atLeast)
}

func compareN(t testing.TB, want, got int, atLeast bool) {
	switch {
	case atLeast && want > got:
		t.Fatalf("want len(seq) >= %d; got %d", want, got)
	case !atLeast && want > got:
		t.Fatalf("want len(seq) == %d; got %d", want, got)
	case !atLeast && want < got:
		t.Fatalf("want len(seq) == %d; got at least %d", want, got)
	}
}

// AtLeastLength calls t.Fatalf if seq has a length that is not at least want.
// The type of seq must be array, array pointer, slice, map, string, channel, [iter.Seq], or [iter.Seq2].
// For channels and iterators, the values are consumed to get the sequence length.
func AtLeastLength(t testing.TB, want int, seq any) {
	t.Helper()
	getLen(t, seq, want, true)
}
