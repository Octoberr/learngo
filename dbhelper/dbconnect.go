package dbhelper

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "Ad123@546db"
	ip       = "h128.ad99.cc"
	port     = "13106"
	dbName   = "dlcms_hospital"
)

func InitDB() {
	path := "root:Ad123@546db@tcp(h128.ad99.cc:13106)/dlcms_hospital"
	DB, _ := sql.Open("mysql", path)
	//defer DB.Close()
	//设置数据库的最大连接
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	err := DB.Ping()
	if err != nil {
		fmt.Println("Open database fail")
		return
	}
	fmt.Println("Connect success")
}
