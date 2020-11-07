package exec

import (
	"regexp"
	"testing"
)

var (
	csv_test_handler = &Handler{
		NumMatch: regexp.MustCompile(`^[0-9]+$`),
	}
	test_records_pattern1 = [][]string{
		[]string{"age", "name", "message"},
		[]string{"25", "Smith", "Hello, everyone."},
		[]string{"30", "Bob", "Glad to meet you."},
	}
	test_records_pattern2 = [][]string{
		[]string{"age", "name", "message"},
	}
	test_records_pattern3 = [][]string{
		[]string{"age", "", "message"},
		[]string{"25", "Smith", "Hello, everyone."},
		[]string{"30", "Bob", "Glad to meet you."},
	}
	test_records_pattern4 = [][]string{
		[]string{"age", "id", "message"},
		[]string{"25", "Smith", "Hello, everyone."},
		[]string{"30", "Bob", "Glad to meet you."},
	}
	expected = map[int]string{
		0: "int",
		1: "varchar(255)",
		2: "varchar(255)",
	}
)

func TestDecideColumnType(t *testing.T) {
	res, err := csv_test_handler.decideColumnType(test_records_pattern1)
	if !isDataTypeMapEqual(res, expected) || err != nil {
		t.Errorf("Pattern1 failed. res: %v, err: %s", res, err)
	}

	res, err = csv_test_handler.decideColumnType(test_records_pattern2)
	if res != nil || err.Error() != "CSV must consist of at least 2 lines" {
		t.Errorf("Pattern2 failed. res: %v, err: %s", res, err)
	}

	res, err = csv_test_handler.decideColumnType(test_records_pattern3)
	if res != nil || err.Error() != "Column names can not be empty" {
		t.Errorf("Pattern3 failed. res: %v, err: %s", res, err)
	}

	res, err = csv_test_handler.decideColumnType(test_records_pattern4)
	if res != nil || err.Error() != "\"id\" can not be used as a column name" {
		t.Errorf("Pattern4 failed. res: %v, err: %s", res, err)
	}
}

func isDataTypeMapEqual(map1, map2 map[int]string) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}
	return true
}
