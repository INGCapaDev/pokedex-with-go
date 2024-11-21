package main

import "fmt"

func cmdPokedex(cfg *config, _ ...string) error {
	fmt.Println("\n Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("-", pokemon.Name)
	}
	fmt.Println()
	return nil
}
