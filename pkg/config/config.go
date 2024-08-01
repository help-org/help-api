package config

import (
	"time"
)

type Config struct {
	Debug bool `envconfig:"DEBUG" default:"false"`

	Server struct {
		Address            string        `envconfig:"SERVER_ADDRESS" default:":8080"`
		ReadTimeout        time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"5s"`
		ReadHeaderTimeout  time.Duration `envconfig:"SERVER_READ_HEADER_TIMEOUT" default:"5s"`
		WriteTimeout       time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"10s"`
		IdleTimeout        time.Duration `envconfig:"SERVER_IDLE_TIMEOUT" default:"120s"`
		MaxHeaderBytes     int           `envconfig:"SERVER_MAX_HEADER_BYTES" default:"120"`
		CorsAllowedOrigins []string      `envconfig:"SERVER_CORS_ALLOWED_ORIGINS" default:"*"`
	}

	Database struct {
		Driver string `envconfig:"DATABASE_DRIVER" default:"sqlite3"`
		Source string `envconfig:"DATABASE_SOURCE" default:"file::memory:?cache=shared"`
	}
}
