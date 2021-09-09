// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
	User   *User  `json:"user"`
}

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (s *User) GenerateHashedPassword() error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(s.Password), 4)
	if err != nil {
		return fmt.Errorf("erorr occured while hashing pass:%v", err)
	}
	s.Password = string(passHash)
	return nil
}

type Userinput struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}