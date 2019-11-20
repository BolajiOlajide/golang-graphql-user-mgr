package golang_graphql_user_mgr

import (
	"context"
	"errors"
	"strconv"
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
		CreatedAt: nanos,
		UpdatedAt: 0,
	}

	err := user.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	err = user.Create()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*Token, error) {
	user, err := FindUserByEmail(email)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	if !user.ComparePassword(password) {
		return nil, errors.New("email or password isn't correct")
	}

	expiredAt := time.Now().Add(time.Hour * 1).Unix()

	userID, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	token := &Token{
		Token:     JWTCreate(int(userID), expiredAt),
		ExpiredAt: int(expiredAt),
	}

	return token, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return GetAllUsers()
}
