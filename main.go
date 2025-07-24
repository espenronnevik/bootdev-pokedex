package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Errorf("Error during input scanning: %w", err)
			break
		}
		inputText := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", inputText[0])
	}
}

// Split string into words, make lowercase and trim whitespace
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
