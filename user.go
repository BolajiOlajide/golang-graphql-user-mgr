package golang_graphql_user_mgr

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// User overridden user model
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Create used to save a user to the database
func (user *User) Create() error {
	result, err := DB.Exec("INSERT INTO `users` (first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);", user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = strconv.Itoa(int(lastID))
	return nil
}

// HashPassword method for hashing user's password
func (user *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// ComparePassword method to validate a user's password
func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
