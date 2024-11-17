package main

import "fmt"

func cmdExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	fmt.Printf("\nExploring %s location...\n", args[0])
	fmt.Printf("Found a wild pokemon! (not implemented yet)\n\n")
	return nil
}
