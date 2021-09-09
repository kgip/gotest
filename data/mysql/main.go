package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/thep0y/go-logger/log"
	"time"
)

type gender byte

const (
	male gender = iota
	female
)

type User struct {
	Id int
	Name string
	Age int
	Gender gender
	Birth time.Time
}

func main() {
	//连接数据库
	db, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/gotest?charset=utf8mb4&parseTime=true&loc=Local")
	defer db.Close()
	if err != nil {
		log.Error(err)
	} else {
		result, err := db.Exec("insert into user(name,age,gender,birth) values(?,?,?,?)", "fusheng",19,0,"2001-10-3")
		if err != nil {
			log.Error(err)
		} else {
			log.Info(result.LastInsertId())
		}
		result, err = db.Exec("update user set age = 21 where id = ?", 1)
		if err != nil {
			log.Error(err)
		} else {
			log.Info(result.RowsAffected())
		}
		users := make([]*User,0,5)
		rows, err := db.Query("select * from user")
		for rows.Next() {
			id := 0
			name := ""
			age := 0
			gend := male
			birth := time.Now()
			//user := []interface{}{&id,&name,&age,&gender,&birth}
			err := rows.Scan(&id,&name,&age,&gend,&birth)
			if err != nil {
				log.Error(err)
			} else {
				log.Info(id,name,age,gend,birth)
				users = append(users, &User{
					id,
					name,
					age,
					gend,
					birth,
				})
			}
		}
		json, err := json.Marshal(users)
		if err != nil {
			log.Info(err)
		} else {
			log.Info(string(json))
		}
		//db.Get("user","select * from user")
	}
}
