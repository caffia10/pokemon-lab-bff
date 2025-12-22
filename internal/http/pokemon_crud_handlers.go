package http

import (
	"pokemon-lab-bff/internal/core"
	"pokemon-lab-bff/internal/http/httpclient"

	"github.com/gofiber/fiber/v3"
)

func CreatePokemonHandler(pkmCli *httpclient.PokemonClient) fiber.Handler {
	return func(c fiber.Ctx) error {

		var pkm core.Pokemon
		if err := c.Bind().Body(&pkm); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if err := pkmCli.CreatePokemon(&pkm); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func GetPokemonHandler(pkmCli *httpclient.PokemonClient) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")

		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing Pokemon ID")
		}

		pkm, err := pkmCli.GetPokemon(id)
		if err != nil {
			return err
		}

		return c.JSON(pkm)
	}
}

func UpdatePokemonHandler(pkmCli *httpclient.PokemonClient) fiber.Handler {
	return func(c fiber.Ctx) error {

		var pkm core.Pokemon
		if err := c.Bind().Body(&pkm); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if err := pkmCli.UpdatePokemon(&pkm); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func DeletePokemonHandler(pkmCli *httpclient.PokemonClient) fiber.Handler {
	return func(c fiber.Ctx) error {

		id := c.Params("id")

		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing Pokemon ID")
		}

		if err := pkmCli.DeletePokemon(id); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func ListPokemonsHandler(pkmCli *httpclient.PokemonClient) fiber.Handler {
	return func(c fiber.Ctx) error {
		pkm, err := pkmCli.ListPokemons()
		if err != nil {
			return err
		}

		return c.JSON(pkm)
	}
}
