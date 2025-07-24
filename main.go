package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var replCommands map[string]cliCommand

func registerReplCommand(name, desc string, cb func() error) error {
	_, exists := replCommands[name]
	if exists {
		return fmt.Errorf("Command already exists")
	}

	replCommands[name] = cliCommand{
		name:        name,
		description: desc,
		callback:    cb,
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range replCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func processCommand(input []string) error {
	name := input[0]
	command, ok := replCommands[name]
	if !ok {
		return fmt.Errorf("Unknown command: %s", name)
	}

	err := command.callback()
	if err != nil {
		return fmt.Errorf("Error executing command: %s\n", command.name)
	}
	return nil
}

func main() {
	var err error
	replCommands = make(map[string]cliCommand)

	err = registerReplCommand("exit", "Exit  the Pokedex", commandExit)
	if err != nil {
		fmt.Println("Error registering 'exit' command.")
	}

	err = registerReplCommand("help", "Display a help message", commandHelp)
	if err != nil {
		fmt.Println("Error registering 'help' command.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if ok {
			inputText := cleanInput(scanner.Text())
			if len(inputText) > 0 {
				err = processCommand(inputText)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
			}
		} else {
			if err = scanner.Err(); err != nil {
				fmt.Printf("Error occured during input scanning: %s\n", err)
				os.Exit(1)
			}
			// EOF on stdin, call exit command
			commandExit()
		}
	}
}

// Split string into words, make lowercase and trim whitespace
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
