package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"xcms/app/model"
	"xcms/app/service"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {

	var err error
	fmt.Println("hello world")
	dsn := "root:123456@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if err := sqlDb.Ping(); err != nil {
		panic(err)
	}
	time.AfterFunc(time.Second*10, func() {
		fmt.Println("10秒后执行")
	})
	admin := new(model.Admin)
	db.Find(admin)
	adminService := new(service.AdminService)
	adminService.Db = db
	list, error := adminService.GetList(1, 2)
	for _, v := range list {
		fmt.Println(v)
	}
	fmt.Printf("%v", list)
	fmt.Println(error)
	fmt.Println(len(list))

	//ticker := time.Tick(time.Second * 5)
	//for i := range ticker {
	//	fmt.Println("5秒执行一次", i)
	//	fmt.Println("admin", admin)
	//	if data, err := json.Marshal(admin); err == nil {
	//		fmt.Println(string(data))
	//	}
	//}
}
