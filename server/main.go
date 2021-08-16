package main

import (
	"log"
	"net/http"

	"github.com/baldore/angelo/db"
	"github.com/baldore/angelo/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dbConn := db.CreateConn()
	queries := db.New(dbConn)

	songsRouter := routes.NewSongsController(dbConn, queries)

	r.Get("/songs", songsRouter.GetSongs)
	r.Post("/songs", songsRouter.CreateSong)

	r.Get("/songs/{id}", songsRouter.GetSong)
	r.Delete("/songs/{id}", songsRouter.DeleteSong)
	r.Patch("/songs/{id}/labels", songsRouter.UpdateLabels)

	log.Fatal(http.ListenAndServe(":4000", r))
}
