package postgres

import (
	"gqlgen-todos/graph/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) GetUsers() []*model.User {
	var users []*model.User
	u.DB.Find(&users)
	return users

}
