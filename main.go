package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"github.com/jaharbaugh/pokedex/internal/pokeapi"
	"github.com/jaharbaugh/pokedex/internal/pokecache"
)

func main() {

	cache := pokecache.NewCache(5 * time.Second)
    apiClient := pokeapi.NewClient("https://pokeapi.co/api/v2", cache)

	cfg := &config{
		next: "",
		previous : "",
		client: apiClient,	
	} 

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0{
			continue
		}

		if cmd, exists := commandList[words[0]]; exists {
			if err := cmd.callback(cfg); err != nil{
				fmt.Println("Error: ", err)
			}
		} else{
				fmt.Println("Unknown command")
		}
	}
}