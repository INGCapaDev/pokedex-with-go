package main

import "fmt"

func cmdInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("\nUsage: %s <pokemon-name>", INSPECT_CMD)
		return fmt.Errorf("you must provide a pokemon name")
	}
	return nil
}
