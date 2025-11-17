package main

import (
	"github.com/jaharbaugh/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name	string
	description	string
	callback	func(*config, []string) error

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
	"explore": {
		name: "explore",
		description: "Displays the pokemon that can be found at a map location",
		callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Throw a pokeball and try to catch a pokemon",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "View information about the pokemon you have captured",
		callback: commandInspect,
	},
	"pokedex": {
		name: "pokedex",
		description: "View a list of all pokemon captured",
		callback: commandPokedex,
	},
	
	}
}