package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string

	ConnStr string
}

func Conn(m *MySqlConfig) (db *sql.DB, err error) {
	connStr := ""
	for {
		if m.ConnStr != "" {
			connStr = m.ConnStr
			break
		}
		if m.Host == "" {
			m.Host = "localhost"
		}
		if m.Port == "" {
			m.Port = "3306"
		}
		if m.User == "" {
			m.User = "root"
		}
		if m.Password == "" {
			m.Password = "root"
		}
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.User, m.Password, m.Host, m.Port, m.Database)
		break
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
}
