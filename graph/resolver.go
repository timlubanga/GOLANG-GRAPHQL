package graph

import (
	"gqlgen-todos/graph/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodosRepo postgres.TodosRepo
	UserRepo  postgres.UserRepo
}
