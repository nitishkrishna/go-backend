package dataset

import (
	"github.com/gocarina/gocsv"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"os"
)

func ReadFromCSV(datasetFile string) {
	// Open the CSV file for reading
	file, err := os.Open(datasetFile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Parse CSV into a slice of typed data `[]GoodreadsBook` (just like json.Unmarshal() does)
	// The builtin package `encoding/csv` does not support unmarshaling into a struct
	// thus, you need to use an external library to avoid writing for-loops.
	var bookEntries []book.GoodreadsBook
	err = gocsv.Unmarshal(file, bookEntries)
	if err != nil {
		panic(err)
	}
}
