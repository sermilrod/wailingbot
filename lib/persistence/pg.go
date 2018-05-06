package pg

import (
	"database/sql"

	config "github.com/sermilrod/wailingbot/lib/configuration"

	// blank import required for postgres driver
	_ "github.com/lib/pq"

	// blank import required for database migrations source
	_ "github.com/mattes/migrate/source/file"
)

// Client returns a postgres client
func Client(cfg *config.Configuration) (*sql.DB, error) {
	connStr := "host=" + cfg.PgHost + " user=" + cfg.PgUser + " password=" + cfg.PgPassword + " dbname=" + cfg.PgDbName + " sslmode=disable connect_timeout=2"
	dbConn, err := sql.Open("postgres", connStr)
	return dbConn, err
}
