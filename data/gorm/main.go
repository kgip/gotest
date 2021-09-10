package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gomod/data/gorm/dao"
	"gomod/data/gorm/entity"
	"math/rand"
	"time"
)

func main() {
	//for i := 0; i < 100; i++ {
	//	age := rand.Intn(20)+10
	//	dao.UserDao.Insert(&entity.User{
	//		Name: uuid.NewV4().String()[:8],
	//		Age: age,
	//		Gender: 1,
	//		Birth: time.Now().AddDate(-age,0,0),
	//	})
	//}
	//for i := 0; i < 100; i++ {
	//	age := rand.Intn(20) + 10
	//	dao.UserDao.Insert(&entity.User{
	//		Name:   uuid.NewV4().String()[:8],
	//		Age:    age,
	//		Gender: 0,
	//		Birth:  time.Now().AddDate(-age, 0, 0),
	//	})
	//}
	age := rand.Intn(20) + 10
	row := dao.UserDao.InsertList([]*entity.User{
		{
			Name:   uuid.NewV4().String()[:8],
			Age:    age,
			Gender: 1,
			Birth:  time.Now().AddDate(-age, 0, 0),
		},
		{
			Name:   uuid.NewV4().String()[:8],
			Age:    age,
			Gender: 1,
			Birth:  time.Now().AddDate(-age, 0, 0),
		},
	})

	fmt.Println(row)
}
