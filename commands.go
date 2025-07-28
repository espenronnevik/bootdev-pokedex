package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/espenronnevik/bootdev-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*state) error
}

type state struct {
	pokedex  map[string]pokeapi.Pokemon
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
	if conf.arg == "" {
		conf.arg = "Goodbye!"
	}
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

func commandExplore(conf *state) error {
	if conf.arg == "" {
		return errors.New("A location name to explore is required")
	}

	fmt.Printf("Exploring %s...\n", conf.arg)
	locarea, err := pokeapi.GetLocationArea(conf.arg)
	if err != nil {
		return fmt.Errorf("Explore command error: %w", err)
	}

	fmt.Println("Found pokemon: ")
	for i := range locarea.PokemonEncounters {
		fmt.Printf(" - %s\n", locarea.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *state) error {
	if conf.arg == "" {
		return errors.New("The name of the pokemon you want to catch is required")
	}

	pokemon, err := pokeapi.GetPokemon(conf.arg)
	if err != nil {
		return fmt.Errorf("Catch command error: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	skill := rand.Intn(700)

	if skill > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func processCommand(input []string, conf *state) error {
	cmdname := input[0]

	// Clear arguments in state after processing completes
	defer func() { conf.arg = "" }()

	if len(input) > 1 {
		conf.arg = input[1]
	}

	command, ok := replCommands[cmdname]
	if !ok {
		return fmt.Errorf("Unknown command: %s", cmdname)
	}

	err := command.callback(conf)
	if err != nil {
		return fmt.Errorf("Error executing command: %s\n", command.name)
	}

	return nil
}
