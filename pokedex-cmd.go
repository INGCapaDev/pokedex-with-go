package main

import (
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

func cmdPokedex(cfg *config.TConfig, _ ...string) error {
	fmt.Println("\n Your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Println("-", pokemon.Name)
	}
	fmt.Println()
	return nil
}
