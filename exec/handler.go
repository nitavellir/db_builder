package exec

import (
	"db_builder/lib/infra/mysql"
)

type Handler struct {
	MysqlDriver *mysql.MysqlDriver
	CSVPath     string
	ErrorMsg    string
}
