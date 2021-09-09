package postgres

import (
	"fmt"
	"gqlgen-todos/graph/model"
	"log"

	"gorm.io/gorm"
)

type TodosRepo struct {
	DB *gorm.DB
}

func (u *TodosRepo) CreateTodo(input model.NewTodo) *model.Todo {
	var user model.User
	u.DB.Take(&user)
	todo := &model.Todo{}
	u.DB.Create(&todo)
	log.Println(*todo)
	return todo

}

func (u *TodosRepo) GetTodos() []*model.Todo {
	var todos []*model.Todo
	u.DB.Find(&todos)
	u.DB.Joins("User").Find(&todos)
	for _, todo := range todos {
		fmt.Println(todo)
	}
	return todos

}

func (u *TodosRepo) GetTodoById() *model.Todo {
	var todo *model.Todo
	u.DB.Find(&todo)
	u.DB.Joins("User").First(&todo)

	return todo

}
