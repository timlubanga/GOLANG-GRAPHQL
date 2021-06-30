package graph

import (
	"gqlgen-todos/graph/model"
	"gqlgen-todos/graph/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Todos    []*model.Todo
	Animals  []*model.Animal
	UserRepo postgres.UserRepo
}
