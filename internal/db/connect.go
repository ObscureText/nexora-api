package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("NEXORA_DB_HOST"),
		os.Getenv("NEXORA_DB_PORT"),
		os.Getenv("NEXORA_DB_USERNAME"),
		os.Getenv("NEXORA_DB_PASSWORD"),
		os.Getenv("NEXORA_DB_NAME"),
		os.Getenv("NEXORA_DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to DB")

	return db, nil
}
