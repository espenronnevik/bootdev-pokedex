package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

var replCommands map[string]cliCommand

func registerReplCommand(name, desc string, cb func(*config) error) error {
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

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range replCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMap(conf *config) error {
	url := urlLocationArea
	if conf.next != "" {
		url = conf.next
	}

	locarea, err := getLocationArea(url)
	if err != nil {
		return fmt.Errorf("Map command error: %w", err)
	}

	if locarea.Next != nil {
		conf.next = *locarea.Next
	} else {
		conf.next = ""
	}

	if locarea.Previous != nil {
		conf.previous = *locarea.Previous
	} else {
		conf.previous = ""
	}

	for _, v := range locarea.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(conf *config) error {
	if conf.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locarea, err := getLocationArea(conf.previous)
	if err != nil {
		return fmt.Errorf("Map command error: %w", err)
	}

	if locarea.Next != nil {
		conf.next = *locarea.Next
	} else {
		conf.next = ""
	}

	if locarea.Previous != nil {
		conf.previous = *locarea.Previous
	} else {
		conf.previous = ""
	}

	for _, v := range locarea.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func processCommand(input []string, conf *config) error {
	name := input[0]
	command, ok := replCommands[name]
	if !ok {
		return fmt.Errorf("Unknown command: %s", name)
	}

	err := command.callback(conf)
	if err != nil {
		return fmt.Errorf("Error executing command: %s\n", command.name)
	}
	return nil
}
