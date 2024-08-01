package cli

import (
	"directory/pkg/config"
	"github.com/kelseyhightower/envconfig"
)

func load() (cfg *config.Config, err error) {
	cfg = new(config.Config)
	err = envconfig.Process("", cfg)
	return
}
