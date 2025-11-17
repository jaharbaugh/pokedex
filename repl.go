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
	for _, i := range pokelist.PokemonEncounters{
		fmt.Println(" -", i.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args[]string) error{
	if len(args) < 1{
		return fmt.Errorf("Invalid target")
	}
	targetPokemon := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + targetPokemon

	target, err := fetchPokemonData(cfg.client, url)
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", targetPokemon)
	CatchAttempt(target)

	return nil
}

func commandInspect(cfg *config, args[]string) error {
	if len(args) < 1{
		return fmt.Errorf("Invalid request")
	}

	pokemonName := args[0]

	target, exists := someonesPC[pokemonName]
    	if !exists {
        	return fmt.Errorf("you have not caught that pokemon")
    	}
	fmt.Printf("Name: %s\n", target.Name)
	fmt.Printf("Height: %d\n", target.Height)
	fmt.Printf("Weight: %d\n", target.Weight)
	fmt.Println("Stats:")
	for _, i := range(target.Stats){
		fmt.Printf("	-%s: %d\n", i.Stat.Name, i.BaseStat)
	}
	fmt.Println("Types:")
	for _, i := range(target.Types){
		fmt.Printf("	- %s\n", i.Type.Name)
	}
	
	return nil
}

func commandPokedex(cfg *config, args[]string) error {
	fmt.Println("Your Pokeddex:")
	for k, _ := range someonesPC{
		fmt.Printf("	- %s\n", k)
	}
	return nil
}