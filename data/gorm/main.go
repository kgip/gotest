package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/thep0y/go-logger/log"
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
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gotest?parseTime=true")
	db.SingularTable(true)
	db.LogMode(true)
	defer db.Close()
	if err != nil {
		log.Error(err)
	}
	users := make([]User, 1, 10)
	//user := &User{}
	db.Where("id > ? or age > ?", 10, 18).Where("name like ?", "%she%").Find(&users)
	log.Info(users)
}
