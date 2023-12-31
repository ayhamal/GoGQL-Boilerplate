package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/ayhamal/gogql-boilerplate/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, userInput model.UserInput) (*model.AuthData, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, email string) (bool, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - resetPassword"))
}

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, user string, content string) (string, error) {
	panic(fmt.Errorf("not implemented: PostMessage - postMessage"))
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.AuthData, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// Messages is the resolver for the messages field.
func (r *subscriptionResolver) Messages(ctx context.Context) (<-chan []*model.Message, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
