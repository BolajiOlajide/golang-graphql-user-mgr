package golang_graphql_user_mgr

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // importing all mysql stuff
)

// DB database connection instance
var DB *sql.DB

// NewDatabase method for creating new database connection
func NewDatabase() {
	databaseURL := os.Getenv("MYSQL_URL")
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		panic(err)
	}

	DB = db
}
