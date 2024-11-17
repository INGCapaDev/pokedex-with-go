package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

type cmd struct {
	name string
	cmd  string
}

const (
	HELP_CMD    = "help"
	EXIT_CMD    = "exit"
	MAP_CMD     = "map"
	MAPB_CMD    = "mapb"
	EXPLORE_CMD = "explore"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startREPL(cfg *config) {
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
	callback    func(cfg *config, args ...string) error
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
			name:        EXPLORE_CMD,
			description: "Explore a location",
			callback:    cmdExplore,
		},
	}
}
