package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type Message struct {
	Message string `json:"message"`
}

func writeJSONMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(Message{message}); err != nil {
		log.Printf("error decoding: %v", err)
	}
}

func main() {
	r := chi.NewRouter()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	connStr := "postgres://postgres:asdfasdf@localhost:5432/angelo?sslmode=disable"
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging to database: %v", err)
	}

	r.Get("/songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "message": "hola mundo" }`)
	})

	r.Post("/songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newSong struct {
			Name string
		}
		if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
			log.Printf("error decoding: %v", err)
			writeJSONMessage(w, "error decoding request body", http.StatusBadRequest)

			return
		}

		insertedID := 0
		insertQuery := `
            insert into songs (name) values ($1)
            returning id
        `
		err = db.QueryRow(insertQuery, newSong.Name).Scan(&insertedID)

		if err, ok := err.(*pq.Error); ok {
			log.Printf("error inserting song: %v", err)
			if err.Code.Name() == "unique_violation" {
				writeJSONMessage(w, "song already exists", http.StatusConflict)
			} else {
				writeJSONMessage(w, "error inserting song. Try again later", http.StatusInternalServerError)
			}

			return
		}

		writeJSONMessage(w, fmt.Sprintf("created song with id: %d", insertedID), http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":4000", r))
}
