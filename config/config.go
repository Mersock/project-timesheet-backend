package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	//Config -.
	Config struct {
		APP   `yaml:"app"`
		HTTP  `yaml:"http"`
		MYSQL `yaml:"mysql"`
		LOG   `yaml:"logger"`
	}

	//APP -.
	APP struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	//HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	//MYSQL -.
	MYSQL struct {
		URL string `env-required:"true" yaml:"url" env:"MYSQL_URL"`
	}

	//LOG -.
	LOG struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
