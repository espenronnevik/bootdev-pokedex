package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/espenronnevik/bootdev-pokedex/internal/pokeapi"
)

func main() {
	conf := state{}

	conf.pokedex = make(map[string]pokeapi.Pokemon)
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

	if err := registerReplCommand("catch", "Attempt to catch a pokemon", commandCatch); err != nil {
		fmt.Println("Error registering 'catch' command.")
	}

	if err := registerReplCommand("inspect", "Inspect a pokemon in the pokedex", commandInspect); err != nil {
		fmt.Println("Error registering 'inspect' command.")
	}

	if err := registerReplCommand("pokedex", "Show the pokemons in your pokedex", commandPokedex); err != nil {
		fmt.Println("Error registering 'pokedex' command.")
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
			if err := processCommand(inputText, &conf); err != nil {
				fmt.Printf("%s\n", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		// Something went wrong, report error and exit with errorcode
		conf.code = 1
		conf.arg = fmt.Sprintf("error occured during input scanning: %s\n", err)
	}

	// EOF on stdin, exit gracefully
	conf.arg = "EOF reached. Goodbye!"
	commandExit(&conf)
}

// Split string into words, make lowercase and trim whitespace
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
