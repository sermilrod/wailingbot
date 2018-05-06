package config

import (
	"github.com/caarlos0/env"
)

// Configuration struct stores the configuration of the bot
type Configuration struct {
	Port       string `env:"WW_PORT" envDefault:"3001"`
	SlackToken string `env:"WW_SLACK_TOKEN"`
	PgUser     string `env:"WW_PG_USER" envDefault:"postgres"`
	PgDbName   string `env:"WW_PG_DBNAME" envDefault:"wailing"`
	PgPassword string `env:"WW_PG_PASSWORD" envDefault:"postgres"`
	PgHost     string `env:"WW_PG_HOST" envDefault:"localhost"`
	AddCmd     string `env:"WW_ADD_CMD" envDefault:"/wwadd"`
	GetCmd     string `env:"WW_GET_CMD" envDefault:"/wwget"`
}

// Parse returns the parsed configuration from envvars
func Parse(cfg *Configuration) (err error) {
	err = env.Parse(cfg)
	if err != nil {
		return err
	}
	return nil
}
