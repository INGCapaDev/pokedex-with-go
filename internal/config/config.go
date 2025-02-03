package config

import (
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

const (
	CONFIG_LOCATION_FOLDER = ".pokecli"
	CAUGHT_POKEMON_FILE    = "caught_pokemon.json"
	INVENTORY_FILE         = "inventory.json"
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
		Inventory:        Inventory{Items: make(map[string]InventoryItem), Balance: INITIAL_BALANCE, MaxCapacity: INITIAL_CAPACITY},
		Shop:             *NewShop(),
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

	fmt.Println("PokeCLI: Loading inventory from disk...")
	err = loadInventoryFromDisk(cfg)
	if err != nil {
		return err
	}

	return nil
}
