package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvConfig struct {
	Port    string `env:"PORT" env-default:"8080"`
	Charset string `env:"CHARSET" env-default:"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}

func LoadEnv() *EnvConfig {
	var cfg EnvConfig

	if err := cleanenv.ReadConfig("./.env", &cfg); err != nil {
		slog.Error("Error loading env", "error", err)
	}
	return &cfg
}