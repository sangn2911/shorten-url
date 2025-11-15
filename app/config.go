package app

import (
	"shorten-url/package/database/mysql"
	"shorten-url/package/httputils"
	"shorten-url/package/logger"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
)

type Config struct {
	ServerConfig      `json:"ServiceConfig"`
	mysql.MySQLConfig `json:"MySQLConfig"`
}

type ServerConfig struct {
	Service            string        `env:"SERVICE"`
	Host               string        `env:"HOST"`
	Port               string        `env:"PORT" envDefault:"8080"`
	WriteTimeout       time.Duration `env:"WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout        time.Duration `env:"READ_TIMEOUT" envDefault:"15s"`
	LogLevel           zerolog.Level `env:"LOG_LEVEL" envDefault:"1"`
	PublicDomain       string        `env:"PUBLIC_DOMAIN" envDefault:"localhost:8080"`
	MinShortenIdLength int           `env:"MIN_SHORTEN_IN_LENGTH" envDefault:"6"`
	httputils.CORSConfig
}

func GetConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		logger.Error(err, "fail to parse env variables")
	}
	return cfg
}
