package main

import (
	"embed"
	"fmt"
	"os"

	tmbjsonparse "github.com/jaskian/tmb-tier-site/tmb-json-parse"
	webbuilder "github.com/jaskian/tmb-tier-site/web-builder"
)

const CURRENT_PHASE = 3

var (
	//go:embed data/character-json.json
	example      []byte
	importKeeper embed.FS
)

func main() {
	data, err := tmbjsonparse.ConvertTMBData(example)
	if err != nil {
		panic(err)
	}
	r, err := webbuilder.NewSiteRenderer()
	if err != nil {
		panic(err)
	}

	webPages, err := r.BuildWebsite(data, CURRENT_PHASE)
	if err != nil {
		panic(err)
	}

	for name, page := range webPages {
		fileName := fmt.Sprintf("../docs/%s", name)
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		file.WriteString(page)
	}
}
