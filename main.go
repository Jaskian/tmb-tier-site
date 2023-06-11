package main

import (
	"embed"

	tmbjsonparse "github.com/jaskian/tmb-tier-site/tmb-json-parse"
)

var (
	//go:embed data/tmbdata.json
	example      []byte
	importKeeper embed.FS
)

func main() {
	data, err := tmbjsonparse.ConvertTMBData(example)

}
