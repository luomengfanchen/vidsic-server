package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接的句柄
var Db *sql.DB

// 打开到数据库的连接
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@/vidsic")
	if err != nil {
		fmt.Println(err.Error())
	}
}