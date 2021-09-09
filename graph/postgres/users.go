package postgres

import (
	"fmt"
	"gqlgen-todos/graph/model"

	uuid "github.com/satori/go.uuid"
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

func (u *UserRepo) GetUserById(id string) *model.User {
	var user model.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return nil
	}
	return &user
}

func (u *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) RegisterUser(userDetails model.Userinput) (*model.User, error) {
	userid := fmt.Sprintf("%v", uuid.NewV4())
	var user = model.User{
		ID:        userid,
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}
	err := user.GenerateHashedPassword()
	if err != nil {
		return nil, err
	}
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
