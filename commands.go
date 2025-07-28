package main

import (
	"fmt"
	"os"

	"github.com/espenronnevik/bootdev-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*state) error
}

type state struct {
	code     int
	arg      string
	nextpage string
	prevpage string
}

var replCommands map[string]cliCommand

func registerReplCommand(name, desc string, cb func(*state) error) error {
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

func commandExit(conf *state) error {
	fmt.Printf("Closing the Pokedex... %s\n", conf.arg)
	os.Exit(conf.code)
	return nil
}

func commandHelp(conf *state) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range replCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMap(conf *state) error {
	locarea, err := pokeapi.GetLocationAreaPage(conf.nextpage)
	if err != nil {
		return fmt.Errorf("Map command error: %w", err)
	}

	if locarea.Next != nil {
		conf.nextpage = *locarea.Next
	} else {
		conf.nextpage = ""
	}

	if locarea.Previous != nil {
		conf.prevpage = *locarea.Previous
	} else {
		conf.prevpage = ""
	}

	for _, v := range locarea.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(conf *state) error {
	if conf.prevpage == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locarea, err := pokeapi.GetLocationAreaPage(conf.prevpage)
	if err != nil {
		return fmt.Errorf("Map command error: %w", err)
	}

	if locarea.Next != nil {
		conf.nextpage = *locarea.Next
	} else {
		conf.nextpage = ""
	}

	if locarea.Previous != nil {
		conf.prevpage = *locarea.Previous
	} else {
		conf.prevpage = ""
	}

	for _, v := range locarea.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func processCommand(input []string, conf *state) error {
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
