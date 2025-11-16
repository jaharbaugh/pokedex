package main

import(
	"fmt"
	"encoding/json"
	"github.com/jaharbaugh/pokedex/internal/pokeapi"
)

type locationAreaList struct{
	Count	int
	Next	*string
	Previous *string
	Results	[]struct{
		Name 	string
		URL		string
	}
}

func fetchLocationAreaPage(c *pokeapi.Client, url string) (locationAreaList, error){
	
	body, err := c.GetBytes(url)
	
	if err != nil{
		return locationAreaList{}, err
	}

	var page locationAreaList

	if err := json.Unmarshal(body, &page); err !=nil{
		return locationAreaList{}, err
	}

	return page, nil
} 

func printAndUpdate(cfg *config, page locationAreaList) {
	for _, r := range page.Results {
		fmt.Println(r.Name)
	}

	if page.Next != nil {
		cfg.next = *page.Next
	} else {
		cfg.next = ""
	}
	if page.Previous != nil {
		cfg.previous = *page.Previous
	} else {
		cfg.previous = ""
	}
}
