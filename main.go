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
    	tokens := cleanInput(input)
    	if len(tokens) == 0 {
        	continue
    	}

    	name := tokens[0]
    	args := tokens[1:]

    	cmd, exists := commandList[name]
    	if !exists {
        	fmt.Println("Unknown command")
        	continue
    	}

    	if err := cmd.callback(cfg, args); err != nil {
        	fmt.Println("Error:", err)
    }
}
}