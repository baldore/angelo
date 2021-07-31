package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateDBConnection() *sql.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging to database: %v", err)
	}

	return db
}
