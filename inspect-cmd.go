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

	fmt.Printf("\nInspecting %s... 🔎\n", pokemon.Name)
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")

	for _, pType := range pokemon.Types {
		fmt.Println("-", pType.Type.Name)
	}

	fmt.Println()

	return nil
}
