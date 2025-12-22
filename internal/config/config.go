package config

type Config struct {
	ServerPort    string
	PokemonApiURL string
}

func LoadConfig() *Config {
	// TODO: Load configuration from environment variables or a config file
	return &Config{
		ServerPort:    "8080",
		PokemonApiURL: "http://example.com/api/pokemons/",
	}
}
