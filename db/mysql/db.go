package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ckeyer/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type DBWrapper struct {
	DB   *sql.DB
	Name string
}

type MysqlConfig struct {
	Host     string
	Port     int
	Database string
	UserName string
	Password string
}

func WrapDB(db *sql.DB) (dbWrapper DBWrapper, err error) {
	curDb := ""
	err = db.QueryRow("select database()").Scan(&curDb)
	if err != nil {
		return
	}
	return DBWrapper{
		Name: curDb,
		DB:   db,
	}, nil
}

func ConnectMysqlDB(host string, port int, database, username, password string) DBWrapper {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		username,
		password,
		host,
		port,
		database,
		"utf8")
	return ConnectMysqlDBConnStr(connStr, database)
}

func ConnectMysqlDBByConfig(config MysqlConfig) DBWrapper {
	return ConnectMysqlDB(config.Host, config.Port, config.Database, config.UserName, config.Password)
}

// connStr: tcp:<host>:<port>*<db>/<user>/<pwd>
// name: name of this connection.
func ConnectMysqlDBConnStr(connStr string, name string) DBWrapper {
	dbWrapper := DBWrapper{}
	dbWrapper.Name = name
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		dbWrapper.logError(err, connStr)
	} else {
		dbWrapper.DB = db
	}
	return dbWrapper
}

func (db DBWrapper) logError(err error, sqlText string, args ...interface{}) {
	var errMsg string
	if err != nil {
		logrus.Error("logError", err)
		errMsg = err.Error()
	}
	logrus.Errorf("DB error, db name: (%s), error: (%s), sql: (%s), args: (%v)", db.Name, errMsg, sqlText, args)
}

// TODO implement the QueryXXXErr()

func (db DBWrapper) Query(sqlText string, args ...interface{}) (qr QueryResult) {
	rows, err := db.DB.Query(sqlText, args...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		db.logError(err, sqlText, args)
		qr.Error = err
		return
	}

	qr.Cols, err = rows.Columns()
	if err != nil {
		db.logError(err, sqlText, args)
		qr.Error = err
		return
	}
	qr.Rows = make([]Row, 0, 10) // Default len 10
	for rows.Next() {
		row := make([]interface{}, len(qr.Cols))
		for i, _ := range row {
			row[i] = new([]byte)
		}

		rows.Scan(row...)
		qr.Rows = append(qr.Rows, row)
	}
	return
}

func (db DBWrapper) QueryScalar(sqlText string, args ...interface{}) interface{} {
	rows, err := db.DB.Query(sqlText, args...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		db.logError(err, sqlText, args)
		return nil
	}

	var v interface{}
	if rows.Next() {
		err = rows.Scan(&v)
	}
	if err != nil {
		db.logError(err, sqlText, args)
	}
	return v
}

func (db DBWrapper) QueryScalarStr(sqlText string, args ...interface{}) string {
	v := db.QueryScalar(sqlText, args...)
	if v == nil {
		return ""
	} else {
		return Row([]interface{}{v}).Str(0)
	}
}

func (db DBWrapper) QueryScalarInt(sqlText string, args ...interface{}) int {
	v := db.QueryScalar(sqlText, args...)
	if v == nil {
		return 0
	} else {
		return Row([]interface{}{v}).ForceInt(0)
	}
}

func (db DBWrapper) QueryScalarFloat(sqlText string, args ...interface{}) float64 {
	v := db.QueryScalar(sqlText, args...)
	if v == nil {
		return 0
	} else {
		return Row([]interface{}{v}).ForceFloat(0)
	}
}

func (db DBWrapper) QueryScalarBytes(sqlText string, args ...interface{}) []byte {
	v := db.QueryScalar(sqlText, args...)
	if v == nil {
		return nil
	} else {
		return Row([]interface{}{v}).Bin(0)
	}
}

func (db DBWrapper) QueryScalarBool(sqlText string, args ...interface{}) bool {
	v := db.QueryScalar(sqlText, args...)
	if v == nil {
		return false
	} else {
		return Row([]interface{}{v}).Bool(0)
	}
}

func (db DBWrapper) Exec(sqlText string, args ...interface{}) sql.Result {
	res, err := db.DB.Exec(sqlText, args...)
	if err != nil {
		db.logError(err, sqlText, args)
		return nil
	}
	return res
}

func (db DBWrapper) ExecErr(sqlText string, args ...interface{}) (sql.Result, error) {
	res, err := db.DB.Exec(sqlText, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
