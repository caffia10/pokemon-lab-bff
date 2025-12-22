package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pokemon-lab-bff/internal/config"
	"pokemon-lab-bff/internal/core"
)

type PokemonClient struct {
	hCli *http.Client
	cfg  *config.Config
}

func NewPokemonClient(hc *http.Client, cfg *config.Config) *PokemonClient {
	return &PokemonClient{hCli: hc, cfg: cfg}
}

func (pc *PokemonClient) GetPokemon(id string) (*core.Pokemon, error) {

	req, err := http.NewRequest("GET", pc.cfg.PokemonApiURL+id, nil)

	if err != nil {
		return nil, err
	}

	resp, err := pc.hCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to get pokem %s, status code: %d", id, resp.StatusCode)
	}

	var p core.Pokemon
	jsonDecoder := json.NewDecoder(resp.Body)
	if err := jsonDecoder.Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (pc *PokemonClient) CreatePokemon(pkm *core.Pokemon) error {

	body, err := json.Marshal(pkm)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", pc.cfg.PokemonApiURL, bytes.NewReader(body))

	if err != nil {
		return err
	}

	resp, err := pc.hCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create pokemon %s, status code: %d", pkm.Name, resp.StatusCode)
	}

	return nil
}

func (pc *PokemonClient) UpdatePokemon(pkm *core.Pokemon) error {
	body, err := json.Marshal(pkm)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", pc.cfg.PokemonApiURL, bytes.NewReader(body))

	if err != nil {
		return err
	}

	resp, err := pc.hCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update pokemon %s, status code: %d", pkm.ID, resp.StatusCode)
	}

	return nil
}

func (pc *PokemonClient) DeletePokemon(id string) error {
	req, err := http.NewRequest("DELETE", pc.cfg.PokemonApiURL+id, nil)

	if err != nil {
		return err
	}

	resp, err := pc.hCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete pokemon %s, status code: %d", id, resp.StatusCode)
	}

	var p core.Pokemon
	jsonDecoder := json.NewDecoder(resp.Body)
	if err := jsonDecoder.Decode(&p); err != nil {
		return err
	}

	return nil
}

func (pc *PokemonClient) ListPokemons() ([]core.Pokemon, error) {
	req, err := http.NewRequest("GET", pc.cfg.PokemonApiURL, nil)

	if err != nil {
		return nil, err
	}

	resp, err := pc.hCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get pokemon list")
	}

	var pkms []core.Pokemon
	jsonDecoder := json.NewDecoder(resp.Body)
	if err := jsonDecoder.Decode(&pkms); err != nil {
		return nil, err
	}

	return pkms, nil
}
