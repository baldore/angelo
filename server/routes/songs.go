package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/baldore/angelo/db"
	"github.com/baldore/angelo/models"
	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

type SongsController struct {
	db      *sql.DB
	queries *db.Queries
}

func NewSongsController(dbConn *sql.DB, queries *db.Queries) *SongsController {
	return &SongsController{
		db:      dbConn,
		queries: queries,
	}
}

// Returns all the songs.
func (c *SongsController) GetSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	songs, err := c.queries.ListSongs(context.Background())
	if err != nil {
		WriteJSONMessage(w, "error getting songs", http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		log.Printf("error encoding: %v", err)
	}
}

// Gets a song with the specified id.
func (c *SongsController) GetSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error parsing id: %v", err), http.StatusInternalServerError)

		return
	}

	song, err := c.queries.GetSong(context.Background(), int32(id))
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error getting song: %v", err), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(song); err != nil {
		log.Printf("error encoding: %v", err)
	}
}

// Creates a new song.
func (c *SongsController) CreateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSong models.Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		log.Printf("error decoding: %v", err)
		WriteJSONMessage(w, "error decoding request body", http.StatusBadRequest)

		return
	}

	insertedID := 0
	err := c.db.
		QueryRow("insert into songs (name) values ($1) returning id", newSong.Name).
		Scan(&insertedID)

	if err, ok := err.(*pq.Error); ok {
		log.Printf("error inserting song: %v", err)

		if err.Code.Name() == "unique_violation" {
			WriteJSONMessage(w, "song already exists", http.StatusConflict)
		} else {
			WriteJSONMessage(w, "error inserting song. Try again later", http.StatusInternalServerError)
		}

		return
	}

	WriteJSONMessage(w, fmt.Sprintf("created song with id: %d", insertedID), http.StatusOK)
}

// Update labels.
func (c *SongsController) UpdateLabels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error reading body: %d", err), http.StatusInternalServerError)

		return
	}

	// validate input by unmarshalling
	var labels []models.Label
	if err := json.Unmarshal(body, &labels); err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error unmarshalling value: %v", err), http.StatusInternalServerError)

		return
	}

	_, err = c.db.Exec("UPDATE songs SET labels = $1 WHERE id = $2", string(body), id)
	if err != nil {
		log.Printf("error updating songs labels: %v", err)
		WriteJSONMessage(w, "error updating song labels", http.StatusInternalServerError)

		return
	}

	WriteJSONMessage(w, fmt.Sprintf("update labels for song: %s", id), http.StatusOK)
}

// Deletes a song.
func (c *SongsController) DeleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	_, err := c.db.Exec("DELETE FROM songs WHERE id = $1", id)
	if err != nil {
		log.Printf("error deleting song: %v", err)
		WriteJSONMessage(w, fmt.Sprintf("error deleting song with id %q", id), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
