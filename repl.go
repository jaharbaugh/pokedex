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

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commandList{
		fmt.Printf("%s:%s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {
	url := cfg.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	
	} 
	
	page, err := fetchLocationAreaPage(cfg.client, url)
	if err != nil{
		return err
	} 
	printAndUpdate(cfg, page)
    return nil
}

func commandMapB(cfg *config, args []string) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	page, err := fetchLocationAreaPage(cfg.client, cfg.previous)
	if err != nil {
		return err
	}
	printAndUpdate(cfg, page)
    return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1{
		return fmt.Errorf("Invalid location")
	}
	location := args[0]
	url := ("https://pokeapi.co/api/v2/location-area/" + location)
	pokelist, err := fetchLocationPokemon(cfg.client, url)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, e := range pokelist.PokemonEncounters{
		fmt.Println(" -", e.Pokemon.Name)
	}

	return nil
}
