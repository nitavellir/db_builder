package exec

import (
	"db_builder/lib/domain/service/csv"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

func (h *Handler) procCSV() error {
	//reshape csv file path
	if h.CSVPath == "" {
		return errors.New("CSV file path is not specified")
	} else if h.MysqlDriver.Table == "" {
		return errors.New("Table name is not specified")
	}

	if string([]rune(h.CSVPath)[0]) != "/" {
		wd, err := os.Getwd()
		if err != nil {
			return errors.New("Can not get the working directory")
		}
		h.CSVPath = filepath.Join(wd, h.CSVPath)
	}

	records, err := csv.GetRecords(h.CSVPath)
	if err != nil {
		return err
	}
	h.Records = records

	//decide column types
	data_type_map, err := h.decideColumnType(records)
	if err != nil {
		return err
	}
	h.DataTypeMap = data_type_map

	return nil
}

//private

func (h *Handler) decideColumnType(records [][]string) (map[int]string, error) {
	if len(records) < 2 {
		return nil, errors.New("CSV must consist of at least 2 lines")
	}

	//in case column names are empty
	for _, name := range records[0] {
		if name == "" {
			return nil, errors.New("Column names can not be empty")
		}
		if name == "id" {
			return nil, errors.New("\"id\" can not be used as a column name")
		}
	}

	data_type_map := make(map[int]string, len(records[0]))
	for _, record := range records {
		if record[0] == records[0][0] {
			continue
		}

		for index, value := range record {
			is_num := h.NumMatch.MatchString(value)
			if is_num {
				value_int64, _ := strconv.ParseInt(value, 10, 64)
				if value_int64 <= 2147483647 && value_int64 <= -2147483647 {
					data_type_map[index] = "int"
				} else {
					data_type_map[index] = "bigint"
				}
			} else {
				data_type_map[index] = "varchar(255)"
			}
		}
	}

	return data_type_map, nil
}
