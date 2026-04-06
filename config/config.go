package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Store Store
	Addr  string
}

type Store struct {
	DB string `env:"DB,required"`
}

func Load() Config {
	var cfg Config

	LoadDefaults(&cfg)

	err := godotenv.Load()
	if err != nil {
		log.Println("[WARN] .env not found")
	}

	if err := env.Parse(&cfg); err != nil {
		log.Println("[WARN] .env parse failed")
	}

	return cfg
}

func LoadDefaults(cfg *Config) {
	cfg.Store.DB = "postgres"
	cfg.Addr = ":8080"
}
