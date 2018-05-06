// +build unit

package config_test

import (
	"os"
	"testing"

	config "github.com/sermilrod/wailingbot/lib/configuration"
	"github.com/franela/goblin"
)

var configTest = []struct {
	envvar string
	value  string
}{
	{"WW_PORT", "3000"},
	{"WW_SLACK_TOKEN", "abcd1234"},
	{"WW_PG_USER", "pguser"},
	{"WW_PG_DBNAME", "pgdb"},
	{"WW_PG_PASSWORD", "pgpass"},
	{"WW_PG_HOST", "pghost"},
	{"WW_ADD_CMD", "/customwwadd"},
	{"WW_GET_CMD", "/customwwget"},
}

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Parse", func() {
		g.Before(func() {
			for _, tt := range configTest {
				os.Setenv(tt.envvar, tt.value)
			}
		})
		g.It("Should populate Configuration struct from envvars", func() {
			cfg := config.Configuration{}
			config.Parse(&cfg)
			g.Assert(cfg.Port).Equal("3000")
			g.Assert(cfg.SlackToken).Equal("abcd1234")
			g.Assert(cfg.PgUser).Equal("pguser")
			g.Assert(cfg.PgDbName).Equal("pgdb")
			g.Assert(cfg.PgPassword).Equal("pgpass")
			g.Assert(cfg.PgHost).Equal("pghost")
			g.Assert(cfg.AddCmd).Equal("/customwwadd")
			g.Assert(cfg.GetCmd).Equal("/customwwget")
		})
		g.It("Should return error if unable to parse configuration", func() {
			err := config.Parse(nil)
			g.Assert(err != nil).Equal(true)
		})
	})
}
