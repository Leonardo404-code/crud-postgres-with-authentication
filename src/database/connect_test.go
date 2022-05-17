package database

import (
	"crud-postgres/src/config"
	"database/sql"
	"testing"
)

var testConnect *sql.DB

func init() {
	config.LoadDotEnvTests()
	testConnect = Connect()
}

func TestConnect(t *testing.T) {
	t.Run("testing", func(t *testing.T) {
		config.LoadDotEnvTests()
		if testConnect == nil {
			t.Fatalf("Failed to connect to database, got: %v", testConnect)
		}
	})
}
