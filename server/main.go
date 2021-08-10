package main

import (
	"log"
	"net/http"

	"github.com/baldore/angelo/db"
	"github.com/baldore/angelo/routes"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	db := db.CreateDBConnection()

	songsRouter := routes.NewSongsController(db)

	r.Get("/songs", songsRouter.GetSongs)
	r.Post("/songs", songsRouter.CreateSong)

	r.Get("/songs/{id}", songsRouter.GetSong)
	r.Delete("/songs/{id}", songsRouter.DeleteSong)
	r.Patch("/songs/{id}/labels", songsRouter.UpdateLabels)

	log.Fatal(http.ListenAndServe(":4000", r))
}
