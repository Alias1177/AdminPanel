package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	ApiUrl string `env:"API_URL"`
}

func Load() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
