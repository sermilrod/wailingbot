// +build integration

package pg_test

import (
	"database/sql"
	"testing"

	config "github.com/sermilrod/wailingbot/lib/configuration"
	"github.com/sermilrod/wailingbot/lib/persistence"
	"github.com/franela/goblin"
)

var db *sql.DB
var dbErr error

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Client", func() {
		g.BeforeEach(func() {
			cfg := config.Configuration{}
			err := config.Parse(&cfg)
			if err != nil {
				panic(err.Error())
			}
			db, dbErr = pg.Client(&cfg)
		})
		g.AfterEach(func() {
			defer db.Close()
		})
		g.It("Should return a postgres client", func() {
			g.Assert(dbErr).Equal(nil)
		})
	})
}
