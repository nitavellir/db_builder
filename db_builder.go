package main

import (
	"db_builder/exec"
	"db_builder/lib/infra/mysql"
	"flag"
	"log"
	"regexp"
)

func main() {

	//Get args
	h := &exec.Handler{
		MysqlDriver: &mysql.MysqlDriver{},
		NumMatch:    regexp.MustCompile(`^[0-9]+$`),
	}
	flag.StringVar(&h.MysqlDriver.User, "user", "", "User of the database")
	flag.StringVar(&h.MysqlDriver.Password, "password", "", "Password of the database")
	flag.StringVar(&h.MysqlDriver.Protocol, "protocol", "", "Protocol of the database")
	flag.StringVar(&h.MysqlDriver.Host, "host", "", "Host of the database")
	flag.IntVar(&h.MysqlDriver.Port, "port", 0, "Port of the database")
	flag.StringVar(&h.MysqlDriver.DB, "db", "", "Database name")
	flag.StringVar(&h.MysqlDriver.Table, "table", "", "Table name")
	flag.StringVar(&h.CSVPath, "csv", "", "CSV file path")
	flag.Parse()

	if status := h.Execute(); status != 0 {
		log.Fatal(h.ErrorMsg)
	}
}
