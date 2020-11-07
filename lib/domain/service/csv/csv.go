package csv

import (
	"encoding/csv"
	"os"
)

func GetRecords(file_path string) ([][]string, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rcsv := csv.NewReader(file)
	records, err := rcsv.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
