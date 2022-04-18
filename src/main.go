package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "qdm199814411:qdm199814411@tcp(qdm199814411.my3w.com)/qdm199814411_db?charset=utf8")
	if err != nil {
		log.Println("数据库连接错误", err)
	}
}
