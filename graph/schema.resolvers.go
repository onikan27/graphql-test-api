package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/onikan27/graphql-test-api/graph/generated"
	"github.com/onikan27/graphql-test-api/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	user := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: "Akino",
	}
	return user, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	var (
		message = "友達検索"
		friends []*model.User
	)
	print(message)
	friend := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: "Akinoの友達",
	}
	friends = append(friends, friend)
	return friends, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
