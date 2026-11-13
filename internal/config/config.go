package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

const DefaultListenAddr = ":8080"

type Config struct {
	ListenAddr     string `env:"LISTEN_ADDR,required" envDefault:":8080"`
	GinMode        string `env:"GIN_MODE" envDefault:"debug"`
	MediatedHLSURL string `env:"MED_HLS_URL,required,notEmpty" envDefault:"http://localhost:8888"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return cfg, nil
}
