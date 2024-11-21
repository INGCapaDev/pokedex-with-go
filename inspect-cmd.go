package main

import "fmt"

func cmdInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("\nUsage: %s <pokemon-name>", INSPECT_CMD)
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon yet")
	}

	fmt.Printf("\nInspecting %s... 🔎", pokemon.Name)
	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")

	for _, pType := range pokemon.Types {
		fmt.Printf("- %s\n", pType.Type.Name)
	}

	fmt.Println()

	return nil
}
