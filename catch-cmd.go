package main

import (
	"fmt"
	"math/rand"
)

func cmdCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("\nUsage: %s <pokemon-name>", CATCH_CMD)
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonResponse, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	if poke, exists := cfg.caughPokemon[pokemonResponse.Name]; exists {
		fmt.Printf("you already captured a %s\n\n", poke.Name)
		return nil
	}

	fmt.Printf("\n Throwing a pokeball at %s...", pokemonResponse.Name)
	pokeResistance := rand.Intn(pokemonResponse.BaseExperience + 1)
	if pokeResistance > 40 {
		fmt.Printf("\n😢 Unfortunately %s has escaped, Try again in a few moment!\n\n", pokemonResponse.Name)
		return nil
	}

	fmt.Printf("\n🎉 Congrats! you captured a wild %s", pokemonResponse.Name)
	cfg.caughPokemon[pokemonResponse.Name] = pokemonResponse
	fmt.Printf("\n %s has been registered in your pokedex!\n\n", pokemonResponse.Name)
	return nil
}
