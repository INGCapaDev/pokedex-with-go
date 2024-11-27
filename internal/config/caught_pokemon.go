package config

import (
	"encoding/json"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func loadCaughtPokemonsFromDisk(cfg *TConfig) error {
	fileData, err, notFound := readDataFromDisk(CAUGHT_POKEMON_FILE)
	if err != nil {
		if notFound {
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

	err := writeDataToDisk(cfg.CaughtPokemon, CAUGHT_POKEMON_FILE)
	if err != nil {
		return err
	}

	return nil
}
