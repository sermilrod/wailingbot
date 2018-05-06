package migrations

import (
	"database/sql"
	"log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
)

var migrationsPath = "file://../../migrations"

// Run handles the database migrations
func Run(db *sql.DB) {
	driver, dbErr := postgres.WithInstance(db, &postgres.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	m.Steps(2)
}
