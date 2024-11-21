package main

import (
	"fmt"
	"time"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func main() {
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
	fmt.Println("Welcome to pokecli! A simple CLI for Pokemon data")
	fmt.Println("Developed by: INGCapaDev with ❤️ ")
	fmt.Println("\nType 'help' for help or 'exit' to exit the program")
}
