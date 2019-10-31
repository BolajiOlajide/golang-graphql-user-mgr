package golang_graphql_user_mgr

import (
	"context"
	"time"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	// panic("not implemented")
	// insert into our database

	now := time.Now()
	nanos := now.UnixNano()

	user := &User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: nanos,
		UpdatedAt: nanos,
	}
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	var users []*User

	user1 := &User{
		ID:        "1",
		FirstName: "Bolaji",
		LastName:  "Olajide",
		Email:     "bolaji@olajide.com",
		CreatedAt: 0,
		UpdatedAt: 0,
	}

	user2 := &User{
		ID:        "2",
		FirstName: "Tolulope",
		LastName:  "Duyile",
		Email:     "tolu@duyile.com",
		CreatedAt: 0,
		UpdatedAt: 0,
	}

	users = append(users, user1, user2)
	return users, nil
}
