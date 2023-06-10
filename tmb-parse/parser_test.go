package tmbparse

import (
	"embed"
	"reflect"
	"testing"
)

var (
	//go:embed examples
	exampleFS embed.FS
	examples  = map[string]string{
		"1": "examples/received-export.csv",
	}
)

func TestParse(t *testing.T) {

	t.Run("Parses items", func(t *testing.T) {
		expected := []TMBItem{
			{
				Player:  "Skyquila",
				ItemID:  "45185",
				Offspec: true,
			},
			{
				Player:  "Skyquila",
				ItemID:  "46138",
				Offspec: false,
			},
			{
				Player:  "Zigsim",
				ItemID:  "40637",
				Offspec: false,
			},
		}

		results, err := ParseFile(exampleFS, examples["1"])

		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected %v, got %v", expected, results)
		}
	})
}
