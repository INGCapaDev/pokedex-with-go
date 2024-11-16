package main

import (
	"fmt"
	"time"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second * 5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	welcome()
	startREPL(cfg)
}

func welcome() {
	fmt.Println("Welcome to pokecli! A simple CLI for Pokemon data")
	fmt.Println("Developed by: INGCapaDev with ❤️ ")
	fmt.Println("\nType 'help' for help or 'exit' to exit the program")
}
