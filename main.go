package main

import (
	"fmt"
	"time"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func main() {
	fmt.Println("PokeCLI: Loading your data...")
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg, err := config.GetConfig(pokeClient)
	if err != nil {
		// Panic if there is an error reading the configuration file
		panic(err)
	}

	welcome()
	startREPL(cfg)
}

func welcome() {
	// clear the screen before showing the welcome message
	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome to PokeCLI! A simple CLI pokemon game")
	fmt.Println("Developed by: INGCapaDev with ❤️ ")
	fmt.Println("\nType 'help' for help or 'exit' to exit the program")
}
