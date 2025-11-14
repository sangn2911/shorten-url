package httputils

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type CORSConfig struct {
	AllowOrigins []string `env:"ALLOW_ORIGIN" envDefault:"*"`
	AllowMethods []string `env:"ALLOW_METHODS" envDefault:"GET, POST"`
	AllowHeaders []string `env:"ALLOW_METHODS" envDefault:"*"`
}

func GetCORSMiddleWare(cfg CORSConfig) mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowedOrigins(cfg.AllowOrigins),
		handlers.AllowedOrigins(cfg.AllowMethods),
		handlers.AllowedOrigins(cfg.AllowHeaders),
	)
}
