package graph

import (
	"context"
	"gqlgen-todos/graph/model"
	"reflect"
	"testing"
)

func Test_queryResolver_Todos(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.Todos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.Todos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.Todos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_DeleteAnimal(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.DeleteAnimal
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Animal
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.DeleteAnimal(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.DeleteAnimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.DeleteAnimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
