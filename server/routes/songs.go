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

	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("error decoding: %v", err)
		WriteJSONMessage(w, "error decoding request body", http.StatusBadRequest)

		return
	}

	newSong, err := c.queries.CreateSong(context.Background(), body.Name)
	if err, ok := err.(*pq.Error); ok {
		log.Printf("error inserting song: %v", err)

		if err.Code.Name() == "unique_violation" {
			WriteJSONMessage(w, "song already exists", http.StatusConflict)
		} else {
			WriteJSONMessage(w, "error inserting song. Try again later", http.StatusInternalServerError)
		}

		return
	}

	WriteJSONMessage(w, fmt.Sprintf("created song with id: %d", newSong.ID), http.StatusOK)
}

// Update labels.
func (c *SongsController) UpdateLabels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error converting id: %v", err), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error reading body: %v", err), http.StatusInternalServerError)
		return
	}

	err = c.queries.UpdateSong(context.Background(), db.UpdateSongParams{
		ID:     int32(id),
		Labels: body,
	})
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error updating song: %v", err), http.StatusInternalServerError)
		return
	}

	WriteJSONMessage(w, fmt.Sprintf("update labels for song: %d", id), http.StatusOK)
}

// Deletes a song.
func (c *SongsController) DeleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error converting id: %v", err), http.StatusInternalServerError)
		return
	}

	err = c.queries.DeleteSong(context.Background(), int32(id))
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error deleting song with id %q", id), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
