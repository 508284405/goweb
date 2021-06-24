package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

//全局变量

var (
	Db  *sql.DB //数据库
	err error   //异常信息
)

func init() {
	//1.0 打开一个数据库
	Db, err = sql.Open("mysql", "root:Kibo@2020@tcp(192.168.20.30:13306)/wytest"+"?charset=utf8&loc="+url.QueryEscape("Asia/Shanghai")+"&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//2.0 检验数据库是否连接成功
	err2 := Db.Ping()
	if err2 != nil {
		panic(err2.Error())
	}

	//查询一行
	_ = Db.QueryRow("select * from dict where id = ?", 2)
	//查询多行
	rows, err2 := Db.Query("select * from dict ")
	if err2 != nil {
		fmt.Println("查询多行失败", err2)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	var id int
	var code string
	var value string
	var remark string
	var createTime time.Time
	var updateTime time.Time

	for rows.Next() {
		err2 = rows.Scan(&id, &code, &value, &remark, &createTime, &updateTime)
		if err2 != nil {
			log.Fatal(err)
		}
		log.Println( id, code, value, remark, createTime.Format("2006-01-02 15:04:05"), updateTime)
	}
}
