package main

import (
	"github.com/jaharbaugh/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name	string
	description	string
	callback	func(*config) error

}

type config struct {
	next string
	previous string
	client *pokeapi.Client
}

var commandList map[string]cliCommand

func init(){
	commandList = map[string]cliCommand{
	"exit": {
		name:	"exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},

	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	
	"map": {
		name: "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays the names of the previous 20 location areas",
		callback: commandMapB,
	},

	}
}