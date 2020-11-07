package csv

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	test_file_path   = "test.csv"
	expected_records = [][]string{
		[]string{"age", "name", "message"},
		[]string{"36", "John", "Thank you."},
		[]string{"54", "田中", "こんにちは"},
	}
)

func TestGetRecords(t *testing.T) {
	wd, _ := os.Getwd()
	test_file_path = filepath.Join(wd, test_file_path)

	res, err := GetRecords(test_file_path)
	if err != nil {
		t.Errorf("Err is not nil. err: %s", err)
	}
	if !isRecordsEqual(res, expected_records) {
		t.Errorf("Unexpected records: %v", res)
	}
}

func isRecordsEqual(slice1, slice2 [][]string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index, line := range slice1 {
		for column, value := range line {
			if slice2[index][column] != value {
				return false
			}
		}
	}
	return true
}
