package golang_graphql_user_mgr

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

// FindUserByEmail fetch a user using their email address
func FindUserByEmail(email string) (*User, error) {
	row := DB.QueryRow("SELECT id, first_name, email, last_name, created_at, password, updated_at FROM `users` WHERE email = ?", email)

	var user User

	err := row.Scan(&user.ID, &user.FirstName, &user.Email, &user.LastName, &user.CreatedAt, &user.Password, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
