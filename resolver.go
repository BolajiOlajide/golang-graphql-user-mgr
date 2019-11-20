package golang_graphql_user_mgr

import (
	"context"
	"time"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver resolver for this rubbish
type Resolver struct{}

// Mutation for this rubbish
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query for this rubbish
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	now := time.Now()
	nanos := now.Unix()

	user := &User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: nanos,
		UpdatedAt: 0,
	}

	err := user.Create()
	if err != nil {
		return nil, err
	}

	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return GetAllUsers()
}
