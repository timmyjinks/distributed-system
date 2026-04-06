package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Store Store
}

type Store struct {
	DB string `env:"DB,required"`
}

func Load() Config {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return Config{}
}
