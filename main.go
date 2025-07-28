package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cmdstate := state{}

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

	if err := registerReplCommand("explore", "Explore a specified location area", commandExplore); err != nil {
		fmt.Println("Error registering 'explore' command.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			// Scanner couldn't scan, bail from the loop
			break
		}

		inputText := cleanInput(scanner.Text())
		if len(inputText) > 0 {
			if err := processCommand(inputText, &cmdstate); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		// Something went wrong, report error and exit with errorcode
		cmdstate.code = 1
		cmdstate.arg = fmt.Sprintf("error occured during input scanning: %s\n", err)
	}

	// EOF on stdin, exit gracefully
	cmdstate.arg = "EOF reached. Goodbye!"
	commandExit(&cmdstate)
}

// Split string into words, make lowercase and trim whitespace
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
