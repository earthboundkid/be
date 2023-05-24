package testfile_test

import (
	"testing"

	"github.com/carlmjohnson/be/testfile"
)

func TestJSON(t *testing.T) {
	testfile.Run(t, "testdata/*.json", func(t *testing.T, path string) {
		testfile.EqualJSON(t, path, struct {
			Data any `json:"data"`
		}{
			Data: []struct {
				Field string `json:"field"`
				Value int    `json:"value"`
			}{
				{"foo", 1},
				{"bar", 2},
			},
		})
	})

}
