package api

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql" // <-- Ensure this import!
)

func GetDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	url := fmt.Sprintf("%s?authToken=%s", dbUrl, authToken)

	DB, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil
}
