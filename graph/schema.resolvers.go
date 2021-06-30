package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	// r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateAnimal(ctx context.Context, input model.NewAnimal) (*model.Animal, error) {
	animal := &model.Animal{
		Name:    input.Name,
		Species: input.Species,
	}
	// r.nil = append(r.animals, animal)
	return animal, nil
}

func (r *mutationResolver) DeleteAnimal(ctx context.Context, input model.DeleteAnimal) ([]*model.Animal, error) {
	// res := remove(r.animals, input.Index)
	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return nil, nil
}

func (r *queryResolver) Animals(ctx context.Context) ([]*model.Animal, error) {
	return nil, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserRepo.GetUsers(), nil
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
// func remove(slice []*model.Animal, s int) []*model.Animal {
// 	return append(slice[:s], slice[s+1:]...)
// }
