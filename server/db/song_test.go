package db_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/baldore/angelo/db"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var (
	database *sql.DB
	queries  *db.Queries
)

func TestMain(m *testing.M) {
	var err error

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
			return err
		}

		if err = database.Ping(); err != nil {
			return err
		}

		queries = db.New(database)

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCreateSong(t *testing.T) {
	_, err := queries.CreateSong(context.Background(), "ton doux sourire")
	assert.NoError(t, err, "error creating song: %v", err)
}
