package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	GORM *gorm.DB
)

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	} else {
		GORM = db
	}
}
