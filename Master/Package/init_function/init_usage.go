package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 使用 init 注册驱动
)

var db *sql.DB

func init() {
	// 数据库连接初始化
	var err error
	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
}

func main() {
	// 使用已经初始化好的 db
	defer db.Close()
	// ... 其他代码
}
