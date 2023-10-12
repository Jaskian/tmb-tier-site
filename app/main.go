package main

import (
	"embed"
	"fmt"
	"os"

	tmbjsonparse "github.com/jaskian/tmb-tier-site/tmb-json-parse"
	webbuilder "github.com/jaskian/tmb-tier-site/web-builder"
)

const CURRENT_PHASE = 4

var (
	//go:embed data/character-json.json
	example []byte
	//go:embed data/p1.json
	p1 []byte
	//go:embed data/p2.json
	p2 []byte
	//go:embed data/p3.json
	p3 []byte

	importKeeper embed.FS
)

func main() {

	previousPhaseFiles := []tmbjsonparse.PhaseFile{
		{Phase: 1, File: p1},
		{Phase: 2, File: p2},
		{Phase: 3, File: p3},
	}

	data, err := tmbjsonparse.ConvertTMBData(example, previousPhaseFiles...)
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
