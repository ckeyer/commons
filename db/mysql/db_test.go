package mysql

import (
	"fmt"
	"testing"
)

func TestDBWrapper(t *testing.T) {
	return
	sql := "SELECT * FROM t LIMIT ?"
	dbWrapper := ConnectMysqlDB("192.168.2.230", 3308, "test", "ckeyer", "wangcj")
	res := dbWrapper.Query(sql, 1)
	if res.Cols[0] == "id" {
		t.Fail()
	}
	for _, r := range res.Rows {
		for k, v := range res.Cols {
			fmt.Println(v, ": ", r.Int(k))
		}
	}
}
