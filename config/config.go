package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int `env:"PORT"`
}

func NewConfig() Config {
	godotenv.Load()
	config := Config{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config
}
