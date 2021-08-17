package db_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/baldore/angelo/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var (
	queries  *db.Queries
	database *sql.DB
)

const clearSongsQuery = "DELETE FROM songs;"

// Setups dockertest and migrations for tests.
func TestMain(m *testing.M) {
	const dbName = "angelo"

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("postgres", "9.6", []string{
		"POSTGRES_PASSWORD=secret",
		"POSTGRES_DB=" + dbName,
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		database, err = sql.Open(
			"postgres",
			fmt.Sprintf(
				"postgres://postgres:secret@localhost:%s/%s?sslmode=disable",
				resource.GetPort("5432/tcp"),
				dbName),
		)
		if err != nil {
			return fmt.Errorf("error opening sql connection: %w", err)
		}

		if err = database.Ping(); err != nil {
			return fmt.Errorf("error trying to ping database: %w", err)
		}

		queries = db.New(database)

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// run migrations
	driver, err := postgres.WithInstance(database, &postgres.Config{})

	mig, err := migrate.NewWithDatabaseInstance(
		"file://../sql/migration", "postgres", driver)
	if err != nil {
		log.Fatalf("error creating migration object: %v", err)
	}

	if err = mig.Up(); err != nil {
		log.Fatalf("error running migration.up: %v", err)
	}

	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCreateSong(t *testing.T) {
	const songName = "ton doux sourire"
	song, err := queries.CreateSong(context.Background(), songName)
	assert.NoError(t, err, "error creating song: %v", err)

	assert.NotNil(t, song.ID)
	assert.Equal(t, song.Name, songName)
	assert.Equal(t, song.Labels, json.RawMessage("[]"))
}

func TestListSongs(t *testing.T) {
	ctx := context.Background()

	_, err := database.ExecContext(ctx, clearSongsQuery)
	assert.NoError(t, err)

	song, err := queries.CreateSong(ctx, "foo")
	assert.NoError(t, err, "error creating song: %v", err)

	songs, err := queries.ListSongs(ctx)
	assert.NoError(t, err, "error listing songs: %v", err)

	assert.Equal(t, 1, len(songs))
	assert.Equal(t, song.Name, songs[0].Name)
	assert.Equal(t, song.ID, songs[0].ID)
}

func TestGetSong(t *testing.T) {
	ctx := context.Background()

	newSong, err := queries.CreateSong(ctx, "song to retrieve")
	assert.NoError(t, err, "error creating song: %v", err)

	foundSong, err := queries.GetSong(ctx, newSong.ID)
	assert.NoError(t, err, "error getting song: %v", err)

	assert.Equal(t, newSong.ID, foundSong.ID)
	assert.Equal(t, newSong.Name, foundSong.Name)
}
