package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	conf := config{}
	replCommands = make(map[string]cliCommand)

	if err := registerReplCommand("exit", "Exit  the Pokedex", commandExit); err != nil {
		fmt.Println("Error registering 'exit' command.")
	}

	if err := registerReplCommand("help", "Display a help message", commandHelp); err != nil {
		fmt.Println("Error registering 'help' command.")
	}

	if err := registerReplCommand("map", "Display the next batch of location areas", commandMap); err != nil {
		fmt.Println("Error registering 'map' command.")
	}

	if err := registerReplCommand("mapb", "Display the previous batch of location areas", commandMapb); err != nil {
		fmt.Println("Error registering 'mapb' command.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			// Scanner reported problems, break out of the loop
			break
		}

		inputText := cleanInput(scanner.Text())
		if len(inputText) > 0 {
			if err := processCommand(inputText, &conf); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error occured during input scanning: %s\n", err)
		os.Exit(1)
	}
	// EOF on stdin, call exit command
	commandExit(nil)
}

// Split string into words, make lowercase and trim whitespace
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
