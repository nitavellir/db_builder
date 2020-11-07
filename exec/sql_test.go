package exec

import (
	"db_builder/lib/infra/mysql"
	"testing"
)

var (
	sql_test_handler = &Handler{
		MysqlDriver: &mysql.MysqlDriver{
			Table: "test_table",
		},
		Records: [][]string{
			[]string{"age", "name", "message"},
			[]string{"25", "Smith", "Hi, everyone."},
			[]string{"28", "Alice", "Nice to meet you."},
		},
		DataTypeMap: map[int]string{
			0: "int",
			1: "varchar(255)",
			2: "varchar(255)",
		},
	}
	expected_create_table_schema = "CREATE TABLE test_table (id bigint NOT NULL AUTO_INCREMENT,age int,name varchar(255),message varchar(255),PRIMARY KEY (id));"
	expected_insert_sql          = "INSERT INTO test_table (age,name,message) VALUES ('25','Smith','Hi, everyone.'),('28','Alice','Nice to meet you.');"
)

func TestCreateTableSchema(t *testing.T) {
	err := sql_test_handler.createTableSchema()
	if err != nil {
		t.Errorf("Err is not nil. err: %s", err)
	}
	if sql_test_handler.CreateTableSchema != expected_create_table_schema {
		t.Errorf("Unexpected create table schema: %s", sql_test_handler.CreateTableSchema)
	}
}

func TestCreateInsertSQL(t *testing.T) {
	err := sql_test_handler.createInsertSQL()
	if err != nil {
		t.Errorf("Err is not nil. err: %s", err)
	}
	if sql_test_handler.InsertSQL != expected_insert_sql {
		t.Errorf("Unexpected insert sql: %s", sql_test_handler.InsertSQL)
	}
}
