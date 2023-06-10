package tmbparse

import (
	"encoding/csv"
	"io/fs"
)

type TMBItem struct {
	Player  string
	ItemID  string
	Offspec bool
}

func ParseFile(fs fs.FS, fileName string) (results []TMBItem, err error) {

	file, err := fs.Open(fileName)

	if err != nil {
		return
	}

	csvr := csv.NewReader(file)
	csvr.FieldsPerRecord = -1

	records, err := csvr.ReadAll()

	if err != nil {
		return
	}

	for i := 1; i < len(records); i++ {
		record := records[i]

		results = append(results, TMBItem{
			Player:  record[3],
			ItemID:  record[10],
			Offspec: record[11] == "1",
		})
	}

	return
}
