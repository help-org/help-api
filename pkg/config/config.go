package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug bool `envconfig:"DEBUG" default:"false"`

	Server struct {
		Address            string        `envconfig:"SERVER_ADDRESS"`
		ReadTimeout        time.Duration `envconfig:"SERVER_READ_TIMEOUT"`
		ReadHeaderTimeout  time.Duration `envconfig:"SERVER_READ_HEADER_TIMEOUT"`
		WriteTimeout       time.Duration `envconfig:"SERVER_WRITE_TIMEOUT"`
		IdleTimeout        time.Duration `envconfig:"SERVER_IDLE_TIMEOUT"`
		MaxHeaderBytes     int           `envconfig:"SERVER_MAX_HEADER_BYTES"`
		CorsAllowedOrigins []string      `envconfig:"SERVER_CORS_ALLOWED_ORIGINS"`
	}

	Database struct {
		Driver string `envconfig:"DATABASE_DRIVER"`
		Source string `envconfig:"DATABASE_SOURCE"`

		MaxConnLifetime time.Duration `envconfig:"DATABASE_MAX_CONN_LIFETIME"`
		MaxConnections  int           `envconfig:"DATABASE_MAX_CONNECTIONS"`
		ConnectTimeout  time.Duration `envconfig:"DATABASE_CONNECT_TIMEOUT"`

		ReadTimeout  time.Duration `envconfig:"DATABASE_READ_TIMEOUT"`
		WriteTimeout time.Duration `envconfig:"DATABASE_WRITE_TIMEOUT"`
	}
}

func (cfg *Config) FromEnv() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}
	err = envconfig.Process("", cfg)
	return
}
