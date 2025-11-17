package main


import(
	"encoding/json"
	"github.com/jaharbaugh/pokedex/internal/pokeapi"
)

type locationPokemonList struct {
  PokemonEncounters []struct {
    Pokemon struct {
      Name string `json:"name"`
    } `json:"pokemon"`
  } `json:"pokemon_encounters"`
}


func fetchLocationPokemon(c *pokeapi.Client, url string) (locationPokemonList, error) {

	body, err := c.GetBytes(url)
	
	if err != nil{
		return locationPokemonList{}, err
	}

	var pokelist locationPokemonList

	if err := json.Unmarshal(body, &pokelist); err !=nil{
		return locationPokemonList{}, err
	}

	return pokelist, nil
}