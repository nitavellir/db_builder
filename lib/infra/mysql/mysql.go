package mysql

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlDriver struct {
	User     string
	Password string
	Protocol string
	Host     string
	Port     int
	DB       string
	Table    string
	Conn     *sqlx.DB
}

func (o *MysqlDriver) Init() error {
	if o.DB == "" {
		return errors.New("Specify DB name")
	} else if o.Port < 0 {
		return errors.New("Port can not be less than 0")
	} else if o.User == "" && o.Password != "" {
		return errors.New("Specify password with user")
	}

	dsn := "/" + o.DB
	if o.Host != "" {
		if o.Port > 0 {
			o.Host = fmt.Sprintf("%s:%d", o.Host, o.Port)
		}
		dsn = fmt.Sprintf("(%s)%s", o.Host, dsn)
	}
	if o.Protocol != "" {
		dsn = o.Protocol + dsn
	}
	if o.User != "" {
		if o.Password != "" {
			dsn = fmt.Sprintf("%s:%s@%s", o.User, o.Password, dsn)
		} else {
			dsn = fmt.Sprintf("%s@%s", o.User, dsn)
		}
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	o.Conn = db
	return nil
}

func (o *MysqlDriver) CreateTable(sql string) error {
	conn := o.Conn
	conn.MustExec(sql)
	return nil
}

func (o *MysqlDriver) InsertData(sql string) error {
	conn := o.Conn
	if _, err := conn.Queryx(sql); err != nil {
		return err
	}
	return nil
}
