package be

import (
	"encoding/json"
	"testing"
)

// JSONEquivalent calls t.Fatalf if want and got serialize to different canonical JSON.
func JSONEquivalent(t testing.TB, want string, got any) {
	t.Helper()
	var wantv any
	if err := json.Unmarshal([]byte(want), &wantv); err != nil {
		t.Fatalf("could not canonicalize wanted JSON: %v", err)
	}
	wantb, err := json.MarshalIndent(wantv, "\t", "  ")
	if err != nil {
		t.Fatalf("marshaling want: %v", err)
	}

	gotb, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("marshaling got: %v", err)
	}
	var gotv any
	if err := json.Unmarshal(gotb, &gotv); err != nil {
		t.Fatalf("could not canonicalize got JSON: %v", err)
	}
	gotb, err = json.MarshalIndent(gotv, "\t", "  ")
	if err != nil {
		t.Fatalf("marshaling got: %v", err)
	}
	if string(wantb) != string(gotb) {
		t.Fatalf(
			"want != got;\n\twant:\n\t%s\n\tgot:\n\t%s\n",
			wantb, gotb)
	}
}
