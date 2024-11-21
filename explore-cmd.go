package main

import (
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

func cmdExplore(cfg *config.TConfig, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("\nUsage: %s <location-name>", EXPLORE_CMD)
		return fmt.Errorf("you must provide a location name")
	}

	locationResponse, err := cfg.PokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("\nExploring %s location...\n", args[0])
	if len(locationResponse.PokemonEncounters) < 1 {
		fmt.Printf("Not pokemon found in this area! bad luck!\n\n")
		return nil
	}

	fmt.Printf("\n Found pokemon:")
	for _, pokemon := range locationResponse.PokemonEncounters {
		fmt.Printf("\n - %s", pokemon.Pokemon.Name)
	}

	fmt.Printf("\n\n")
	return nil
}
