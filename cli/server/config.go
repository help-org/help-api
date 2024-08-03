package cli

import (
	"github.com/kelseyhightower/envconfig"

	"directory/pkg/config"
)

func configFromEnv() (cfg *config.Config, err error) {
	cfg = new(config.Config)
	err = envconfig.Process("", cfg)
	return
}
