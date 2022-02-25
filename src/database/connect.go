package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:151452@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return db
}
