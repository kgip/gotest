package main

import (
	"github.com/thep0y/go-logger/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	Id     int `gorm:"PRIMARY_KEY"`
	Name   string
	Age    int
	Gender byte
	Birth  time.Time
}

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	//db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gotest?parseTime=true")
	//sqlDb := db.DB()
	//sqlDb.SetMaxOpenConns(200)
	//sqlDb.SetMaxIdleConns(150)
	//db.SingularTable(true)
	//db.LogMode(true)
	//defer db.Close()
	if err != nil {
		log.Error(err)
	}
	users := make([]User, 1, 10)
	//user := &User{}
	db.Where("id > ? or age > ?", 10, 18).Where("name like ?", "%she%").Find(&users)
	log.Info(users)
}
