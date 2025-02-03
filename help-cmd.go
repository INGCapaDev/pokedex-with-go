package main

import (
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

func cmdHelp(cfg *config.TConfig, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("\nUsage: <command>")
		for _, cmd := range getCommands() {
			fmt.Printf("\n%s: %s", cmd.name, cmd.description)
		}
		fmt.Printf("\n\n")
		return nil
	}
	if len(args) >= 1 {
		switch args[0] {
		case SHOP_CMD:
			return cmdShop(cfg, args[1:]...)
		default:
			{
				fmt.Printf("\nCommand not found. Type '%s' for help or '%s' to exit the program\n\n", HELP_CMD, EXIT_CMD)
			}
		}
	}
	return nil
}
