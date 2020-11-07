package exec

import (
	"fmt"
	"strings"
)

func (h *Handler) createSQL() error {
	//table schema
	if err := h.createTableSchema(); err != nil {
		return err
	}

	//insert SQL
	if err := h.createInsertSQL(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) createTableSchema() error {
	create_table_columns := "id bigint NOT NULL AUTO_INCREMENT,\n"
	for index, record := range h.Records[0] {
		data_type := h.DataTypeMap[index]
		create_table_columns += fmt.Sprintf(`%s %s,\n`, record, data_type)
	}
	create_table_columns = strings.TrimRight(create_table_columns, ",\n")

	create_table_schema := fmt.Sprintf(`
		CREATE TABLE %s (
			%s
			PRIMARY KEY (id)
		);
	`, h.MysqlDriver.Table, create_table_columns)

	h.CreateTableSchema = create_table_schema
	return nil
}

func (h *Handler) createInsertSQL() error {
	//columns
	columns := make([]string, 0, len(h.Records[0]))
	for _, column := range h.Records[0] {
		columns = append(columns, column)
	}
	insert_columns_sql := strings.Join(columns, ",")

	//records
	insert_records := make([]string, 0, len(h.Records)-1)
	for _, record := range h.Records {
		if record[0] == h.Records[0][0] {
			continue
		}

		values := make([]string, 0, len(record))
		for _, value := range record {
			values = append(values, fmt.Sprintf(`\'%s\'`, value))
		}
		one_record := fmt.Sprintf(`(%s)`, strings.Join(values, ","))
		insert_records = append(insert_records, one_record)
	}
	insert_values_sql := strings.Join(insert_records, ",\n")

	insert_sql := fmt.Sprintf(`
		INSERT INTO %s (%s)
		VALUES
		%s;
	`, h.MysqlDriver.Table, insert_columns_sql, insert_values_sql)

	h.InsertSQL = insert_sql
	return nil
}
