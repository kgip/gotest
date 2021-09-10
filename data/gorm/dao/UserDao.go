package dao

import (
	"gomod/data/gorm/config/db"
	"gomod/data/gorm/dao/interfaces"
	"gomod/data/gorm/entity"
)

type userDao struct{}

var (
	UserDao interfaces.IUserDao = &userDao{}
	u                           = &entity.User{}
)

func (userDao) Insert(user *entity.User) int {
	db := db.GORM.Create(user)
	return int(db.RowsAffected)
}

func (userDao) InsertList(users []*entity.User) int {
	db := db.GORM.Create(users)
	return int(db.RowsAffected)
}

func (userDao) DeleteById(id int) int {
	db := db.GORM.Delete(u, id)
	return int(db.RowsAffected)
}

func (userDao) UpdateById(user *entity.User) int {
	db := db.GORM.Find(u, user.ID).Updates(user)
	return int(db.RowsAffected)
}
