package main

import "fmt"

func cmdHelp() error {
	fmt.Printf("\nUsage: <command>")
	for _, cmd := range getCommands() {
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
	}
	fmt.Printf("\n\n")
	return nil
}
