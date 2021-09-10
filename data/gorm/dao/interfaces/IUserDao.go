package interfaces

import (
	"gomod/data/gorm/entity"
)

type IUserDao interface {
	Insert(user *entity.User) int
	InsertList(users []*entity.User) int

	DeleteById(id int) int

	UpdateById(user *entity.User) int
}
