package main

import(
	"encoding/json"
	"github.com/jaharbaugh/pokedex/internal/pokeapi"
	"math/rand"
	"fmt"
)

type Pokemon struct {
    Name           string `json:"name"`
    BaseExperience int    `json:"base_experience"`
	Height	int `json:"height"`
	Weight	int	`json:"weight"`
	Stats	[]PokemonStat `json:"stats"`
	Types []PokemonType `json:"types"`
}

type PokemonStat struct{
	BaseStat int	`json:"base_stat"`
	Stat 	StatInfo 
}

type StatInfo struct{
	Name	string `json:"name"`
}

type PokemonType struct{
	Slot	int		`json:"slot"`
	Type	PokemonTypeInfo		
}

type PokemonTypeInfo struct {
	Name 	string `json:"name"`
}

var someonesPC map[string]Pokemon

func fetchPokemonData(c *pokeapi.Client, url string) (Pokemon, error){
	
	body, err := c.GetBytes(url)
	
	if err != nil{
		var zero Pokemon
		return zero, err
	}

	var p Pokemon
	if err := json.Unmarshal(body, &p); err !=nil{
		return p, err
	}

	return p, nil

}

func CatchAttempt(target Pokemon) {
	if someonesPC == nil {
        someonesPC = make(map[string]Pokemon)
    }	
	randomInt := rand.Intn(608)
	if target.BaseExperience < randomInt{
		fmt.Printf("%s was caught!\n", target.Name)
		someonesPC[target.Name] = target
	} else{
		fmt.Printf("%s escaped!\n", target.Name)
	}
}
