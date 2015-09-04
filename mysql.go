package lib

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

var (
	ms *MySqlConfig
)

func Conn(connStr string) (db *sql.DB, err error) {

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
}

// GetConnStr 获取连接字符串
func GetConnStr() string {
	if ms != nil {
		if ms.ConnStr != "" {
			return ms.ConnStr
		}
		if ms.Host == "" {
			ms.Host = "localhost"
		}
		if ms.Port == "" {
			ms.Port = "3306"
		}
		if ms.User == "" {
			ms.User = "root"
		}
		if ms.Password == "" {
			ms.Password = "root"
		}
		ms.ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.User, m.Password, m.Host, m.Port, m.Database)
		return ms.ConnStr
	}
	return ""
}
