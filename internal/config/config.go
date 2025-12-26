package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort    string `required:"true" split_words:"true"`
	PokemonApiURL string `required:"true" split_words:"true"`
}

func LoadConfig() *Config {
	var cfg Config
	err := envconfig.Process("pokemon-lab-bff", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
