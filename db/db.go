package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PW")
	name := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	DB = db
}
