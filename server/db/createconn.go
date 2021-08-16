package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateConn() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging to database: %v", err)
	}

	return db
}
