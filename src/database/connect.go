package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Connect open connection to database
func Connect() *sql.DB {
	dsn := viper.GetString("POSTGRES_URL")

	if len(dsn) == 0 {
		log.Fatal("missing dsn")
	}

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	db.SetConnMaxLifetime(time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(3)

	err = db.Ping()

	if err != nil {
		log.Fatalf("%v", err)
	}

	db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		password TEXT
	);`)

	return db
}
