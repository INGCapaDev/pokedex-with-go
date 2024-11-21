package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

const (
	CONFIG_LOCATION_FOLDER = ".pokecli"
	CAUGHT_POKEMON_FILE    = "caught_pokemon.json"
)

func GetConfig(pokeClient pokeapi.Client) (*TConfig, error) {
	cfg := newConfig(pokeClient)
	err := cfg.loadDataFromDisk()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func newConfig(pokeClient pokeapi.Client) *TConfig {
	return &TConfig{
		PokeapiClient:    pokeClient,
		CaughtPokemon:    make(map[string]pokeapi.Pokemon),
		NextLocationsURL: nil,
		PrevLocationsURL: nil,
	}
}

func (cfg *TConfig) loadDataFromDisk() error {
	fmt.Println("PokeCLI: Ensuring configuration folder exists...")
	err := ensureConfigFolderExists()
	if err != nil {
		return err
	}

	fmt.Println("PokeCLI: Loading caught pokemons from disk...")
	err = loadCaughtPokemonsFromDisk(cfg)
	if err != nil {
		return err
	}

	return nil
}

func ensureConfigFolderExists() error {
	configLocation, err := getConfigLocation("/")
	if err != nil {
		return err
	}

	if _, err := os.Stat(configLocation); os.IsNotExist(err) {
		err := os.MkdirAll(configLocation, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func getConfigLocation(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename cannot be empty")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configLocation := filepath.Join(homeDir, CONFIG_LOCATION_FOLDER, filename)
	return configLocation, nil
}
