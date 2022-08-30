package dataset

import (
	"github.com/gocarina/gocsv"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"io"
	"os"
	"path/filepath"
)

func ReadFromCSV(datasetFile string) *[]book.GoodreadsBook {
	// Open the CSV file for reading
	pwd, _ := os.Getwd()

	file, err := os.Open(filepath.Join(pwd, datasetFile))
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
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.LazyCSVReader(in)
	})
	err = gocsv.Unmarshal(file, &bookEntries)
	if err != nil {
		panic(err)
	}
	return &bookEntries
}
