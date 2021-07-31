package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/baldore/angelo/db"
	"github.com/go-chi/chi"
	"github.com/lib/pq"
)

type Message struct {
	Message string `json:"message"`
}

type Song struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
	db := db.CreateDBConnection()

	r.Get("/songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		selectSQL := `select id, name from songs`
		rows, err := db.Query(selectSQL)
		if err != nil {
			writeJSONMessage(w, "error getting songs", http.StatusInternalServerError)
			return
		}

		var songs []Song
		for rows.Next() {
			var song Song
			if err := rows.Scan(&song.ID, &song.Name); err != nil {
				log.Printf("error scanning songs row: %v", err.Error())
			}
			songs = append(songs, song)
		}

		if err := json.NewEncoder(w).Encode(songs); err != nil {
			log.Printf("error encoding: %v", err)
		}
	})

	r.Get("/songs/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")
		selectSQL := `select id, name from songs where id = $1`

		var song Song
		err := db.QueryRow(selectSQL, id).Scan(&song.ID, &song.Name)
		if err != nil {
			writeJSONMessage(w, fmt.Sprintf("error querying song with id %s: %v", id, err), http.StatusInternalServerError)

			return
		}

		if err := json.NewEncoder(w).Encode(song); err != nil {
			log.Printf("error encoding: %v", err)
		}
	})

	r.Post("/songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newSong Song
		if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
			log.Printf("error decoding: %v", err)
			writeJSONMessage(w, "error decoding request body", http.StatusBadRequest)

			return
		}

		insertedID := 0
		insertQuery := `insert into songs (name) values ($1) returning id`
		err := db.QueryRow(insertQuery, newSong.Name).Scan(&insertedID)

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
