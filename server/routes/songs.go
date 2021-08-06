package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lib/pq"
)

type Label struct {
	Name string `json:"name"`
}

type Song struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Labels []Label `json:"labels"`
}

type SongsController struct {
	db *sql.DB
}

func NewSongsController(db *sql.DB) *SongsController {
	return &SongsController{
		db,
	}
}

// Returns all the songs
func (c *SongsController) GetSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sqlQuery := `
        select id, name, labels
        from songs
        order by id
    `
	rows, err := c.db.Query(sqlQuery)
	if err != nil {
		WriteJSONMessage(w, "error getting songs", http.StatusInternalServerError)
		return
	}

	var songs []Song
	for rows.Next() {
		var song Song
		var rawLabels string
		if err := rows.Scan(&song.ID, &song.Name, &rawLabels); err != nil {
			log.Printf("error scanning songs row: %v", err.Error())
		}

		var labels []Label
		if err := json.Unmarshal([]byte(rawLabels), &labels); err != nil {
			log.Printf("error unmarshalling value: %v", err.Error())
		}

		song.Labels = labels

		songs = append(songs, song)
	}

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		log.Printf("error encoding: %v", err)
	}
}

// Gets a song with the specified id
func (c *SongsController) GetSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	var song Song
	sqlQuery := `
        select id, name
        from songs
        where id = $1
    `
	if err := c.db.QueryRow(sqlQuery, id).Scan(&song.ID, &song.Name); err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error querying song with id %s: %v", id, err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(song); err != nil {
		log.Printf("error encoding: %v", err)
	}
}

// Creates a new song
func (c *SongsController) CreateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSong Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		log.Printf("error decoding: %v", err)
		WriteJSONMessage(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	insertedID := 0
	err := c.db.QueryRow(
		`insert into songs (name) values ($1) returning id`,
		newSong.Name,
	).Scan(&insertedID)

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

func (c *SongsController) UpdateLabels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error reading body: %d", err), http.StatusInternalServerError)
		return
	}

	// validate input by unmarshalling
	var labels []Label
	if err := json.Unmarshal(body, &labels); err != nil {
		WriteJSONMessage(w, fmt.Sprintf("error unmarshalling value: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = c.db.Exec(
		`update songs
        set labels = $1
        where id = $2`,
		string(body),
		id,
	)
	if err != nil {
		log.Printf("error updating songs labels: %v", err)
		WriteJSONMessage(w, fmt.Sprintf("error updating song labels"), http.StatusInternalServerError)
		return
	}

	WriteJSONMessage(w, fmt.Sprintf("update labels for song: %s", id), http.StatusOK)
}
