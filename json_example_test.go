package be_test

import (
	"github.com/carlmjohnson/be"
)

func ExampleJSONEquivalent() {
	var t mockingT

	// good
	be.JSONEquivalent(t, `{
		"greeting":  "hello",
		"addressee": "world"
	}`, struct {
		Address string `json:"greeting"`
		To      string `json:"addressee"`
	}{"hello", "world"})

	// bad
	be.JSONEquivalent(t, `{
		"Greeting": "hello",
		"Addressee": "world"
	}`, struct {
		Greeting  string
		Addressee string
	}{"salutations", "planet"})
	// Output:
	// want != got;
	// 	want:
	// 	{
	// 	  "Addressee": "world",
	// 	  "Greeting": "hello"
	// 	}
	// 	got:
	// 	{
	// 	  "Addressee": "planet",
	// 	  "Greeting": "salutations"
	// 	}
}
