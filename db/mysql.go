package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}

func Conn(connStr string) (db *sql.DB, err error) {

	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return
}

// GetConnStr 获取连接字符串
func (m *MySqlConfig) GetConnStr() string {
	if m != nil {
		if m.Host == "" {
			m.Host = "localhost"
		}
		if m.Port == 0 {
			m.Port = 3306
		}
		if m.User == "" {
			m.User = "root"
		}
		if m.Password == "" {
			m.Password = "root"
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", m.User, m.Password, m.Host, m.Port, m.Database, m.Charset)
		return connStr
	}
	return ""
}
