package main

import (
	"fmt"
)

func main() {
	welcome()
	startREPL()
}

func welcome() {
	fmt.Println("Welcome to pokecli! A simple CLI for Pokemon data")
	fmt.Println("Developed by: INGCapaDev with ❤️ ")
	fmt.Println("\nType 'help' for help or 'exit' to exit the program")
}
