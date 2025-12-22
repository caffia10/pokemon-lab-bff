package boostrap

import (
	standarHttp "net/http"

	"pokemon-lab-bff/internal/config"
	"pokemon-lab-bff/internal/http"
	"pokemon-lab-bff/internal/http/httpclient"
)

func InitializeApp() {

	httpClient := standarHttp.DefaultClient
	cfg := config.LoadConfig()

	pkmCli := httpclient.NewPokemonClient(
		httpClient,
		cfg,
	)

	httpServer := http.NewAppServer(
		pkmCli,
		cfg,
	)

	httpServer.Start()
}
