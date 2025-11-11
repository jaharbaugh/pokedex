package main

import(
	"strings"
	"os"
	"fmt"
)

func cleanInput(text string) []string {
	cleanText := strings.ToLower(text)
	cleanText = strings.TrimSpace(cleanText)
	return strings.Fields(cleanText)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range commandMap{
		fmt.Printf("%s:%s\n", cmd.name, cmd.description)
	}
	return nil
}