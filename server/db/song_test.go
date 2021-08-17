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

var queries *db.Queries

// Setups dockertest and migrations for tests.
func TestMain(m *testing.M) {
	var database *sql.DB

	const databaseName = "angelo"

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("postgres", "9.6",
		[]string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=" + databaseName})
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
				databaseName))
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
		"file://../sql/migration",
		"postgres", driver)
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
	t.Parallel()

	const songName = "ton doux sourire"
	song, err := queries.CreateSong(context.Background(), songName)

	assert.NoError(t, err, "error creating song: %v", err)

	assert.NotNil(t, song.ID)
	assert.Equal(t, song.Name, songName)
	assert.Equal(t, song.Labels, json.RawMessage("[]"))
}
