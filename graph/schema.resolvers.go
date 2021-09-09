package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := r.TodosRepo.CreateTodo(model.NewTodo{
		Text:   input.Text,
		UserID: input.UserID,
	})
	// r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.Userinput) (*model.AuthToken, error) {
	user, err := r.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if user.Email != "" {
		return nil, fmt.Errorf("user with the email addressed:%s provided exists", input.Email)
	}

	user, err = r.UserRepo.RegisterUser(input)
	if err != nil {
		return nil, err
	}

	token, err := user.GenerateJWTToken()
	if err != nil {
		return nil, err
	}

	return &model.AuthToken{
		User:  user,
		Token: token,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := r.TodosRepo.GetTodos()
	return todos, nil
}

func (r *queryResolver) GetTodoByID(ctx context.Context) (*model.Todo, error) {
	todo := r.TodosRepo.GetTodoById()
	return todo, nil
}

func (r *queryResolver) GetUser(ctx context.Context, input model.Userinput) (*model.User, error) {
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserRepo.GetUsers(), nil
}
