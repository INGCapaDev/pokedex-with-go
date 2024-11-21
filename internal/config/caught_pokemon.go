package config

import (
	"encoding/json"
	"os"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func loadCaughtPokemonsFromDisk(cfg *TConfig) error {
	configLocation, err := getConfigLocation(CAUGHT_POKEMON_FILE)
	if err != nil {
		return err
	}

	fileData, err := os.ReadFile(configLocation)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(fileData, &cfg.CaughtPokemon)
}

func (cfg *TConfig) CatchPokemon(pokemon pokeapi.Pokemon) error {
	if _, exists := cfg.CaughtPokemon[pokemon.Name]; exists {
		return nil
	}

	cfg.CaughtPokemon[pokemon.Name] = pokemon

	data, err := json.MarshalIndent(cfg.CaughtPokemon, "", "  ")
	if err != nil {
		return err
	}

	configLocation, err := getConfigLocation(CAUGHT_POKEMON_FILE)
	if err != nil {
		return err
	}

	err = os.WriteFile(configLocation, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
