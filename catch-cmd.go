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

	if _, exists := cfg.caughPokemon[args[0]]; exists {
		fmt.Printf("you already captured a %s\n\n", args[0])
		return nil
	}

	pokemonResponse, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	pokeResistance := rand.New(rand.NewSource(int64(pokemonResponse.BaseExperience))).Int()

	if pokeResistance < 40 {
		fmt.Printf("\n🎉 Congrats! you captured a wild %s", args[0])
		cfg.caughPokemon[args[0]] = pokemonResponse
		fmt.Printf("\n %s has been registered in your pokedex!\n\n", args[0])
		return nil
	}

	fmt.Printf("\n😢 Unfortunately %s has escaped, Try again in a few moment!\n\n", args[0])
	return nil
}
