package exec

import (
	"db_builder/lib/infra/mysql"
	"regexp"
)

type Handler struct {
	MysqlDriver       *mysql.MysqlDriver
	NumMatch          *regexp.Regexp
	CSVPath           string
	Records           [][]string
	DataTypeMap       map[int]string
	ErrorMsg          string
	CreateTableSchema string
	InsertSQL         string
}
