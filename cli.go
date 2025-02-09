package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

type cmd struct {
	name string
	cmd  string
}

func startREPL(cfg *config.TConfig) {
	for true {
		reader := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter command > ")
		reader.Scan() // Wait for the user input

		if reader.Err() != nil {
			fmt.Println(fmt.Errorf("error reading input: %v", reader.Err()))
			continue
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := getCommands()[words[0]]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if !exists {
			fmt.Printf("\nCommand not found. Type '%s' for help or '%s' to exit the program\n\n", HELP_CMD, EXIT_CMD)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("\nerror executing %s command: %v\n\n", command.name, err)
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config.TConfig, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		HELP_CMD: {
			name:        HELP_CMD,
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		EXIT_CMD: {
			name:        EXIT_CMD,
			description: "Exit the program",
			callback:    cmdExit,
		},
		MAP_CMD: {
			name:        MAP_CMD,
			description: "Show the next page of locations",
			callback:    cmdMap,
		},
		MAPB_CMD: {
			name:        MAPB_CMD,
			description: "Show the previous page of locations",
			callback:    cmdMapB,
		},
		EXPLORE_CMD: {
			name:        EXPLORE_CMD + " <location-name>",
			description: "Explore a location",
			callback:    cmdExplore,
		},
		CATCH_CMD: {
			name:        CATCH_CMD + " <pokemon-name>",
			description: "Make an attempt to catch a wild pokemon",
			callback:    cmdCatch,
		},
		INSPECT_CMD: {
			name:        INSPECT_CMD + " <pokemon-name>",
			description: "Display the basic info of a captured pokemon",
			callback:    cmdInspect,
		},
		POKEDEX_CMD: {
			name:        POKEDEX_CMD,
			description: "Display all the registered pokemon in your pokedex",
			callback:    cmdPokedex,
		},
		SHOP_CMD: {
			name:        SHOP_CMD,
			description: "Type 'help shop' for more information",
			callback:    cmdShop,
		},
	}
}
