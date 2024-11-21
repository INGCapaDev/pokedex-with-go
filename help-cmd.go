package main

import (
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

func cmdHelp(cfg *config.TConfig, args ...string) error {
	fmt.Printf("\nUsage: <command>")
	for _, cmd := range getCommands() {
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
	}
	fmt.Printf("\n\n")
	return nil
}
