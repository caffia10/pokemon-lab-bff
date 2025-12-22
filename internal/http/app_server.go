package http

import (
	"pokemon-lab-bff/internal/config"
	"pokemon-lab-bff/internal/http/httpclient"

	"github.com/gofiber/fiber/v3"
)

type AppServer struct {
	cfg    *config.Config
	router *fiber.App
	pkmCli *httpclient.PokemonClient
}

func NewAppServer(pkmCli *httpclient.PokemonClient, cfg *config.Config) *AppServer {

	router := fiber.New()

	return &AppServer{router: router, pkmCli: pkmCli, cfg: cfg}
}

func (s *AppServer) Start() error {
	//TODO: register other routes/middlewares
	s.registerPokemonCRUDRoutes()
	return s.router.Listen(s.cfg.ServerPort)
}

func (s *AppServer) registerPokemonCRUDRoutes() {
	pokemonCRUDGroup := s.router.Group("/api/pokemons")
	{
		pokemonCRUDGroup.Post("", CreatePokemonHandler(s.pkmCli))
		pokemonCRUDGroup.Get("/:id", GetPokemonHandler(s.pkmCli))
		pokemonCRUDGroup.Put("/:id", UpdatePokemonHandler(s.pkmCli))
		pokemonCRUDGroup.Delete("/:id", DeletePokemonHandler(s.pkmCli))
		pokemonCRUDGroup.Get("", ListPokemonsHandler(s.pkmCli))
	}
}
