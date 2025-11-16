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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range commandList{
		fmt.Printf("%s:%s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	url := cfg.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	
	} 
	
	page, err := fetchLocationAreaPage(url)
	if err != nil{
		return err
	} 
	printAndUpdate(cfg, page)
    return nil
}

func commandMapB (cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	page, err := fetchLocationAreaPage(cfg.previous)
	if err != nil {
		return err
	}
	printAndUpdate(cfg, page)
    return nil

}