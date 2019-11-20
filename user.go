package golang_graphql_user_mgr

import "strconv"

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

// GetAllUsers fetch all users from the database
func GetAllUsers() ([]*User, error) {
	var result []*User

	rows, err := DB.Query("SELECT id, first_name, last_name, email, created_at, updated_at FROM `users`;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &user)
	}

	return result, nil
}
