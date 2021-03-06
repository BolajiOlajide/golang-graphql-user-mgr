// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package golang_graphql_user_mgr

type NewUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Token struct {
	Token     string `json:"token"`
	ExpiredAt int    `json:"expired_at"`
}
