package main

import "strings"
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config,  ...string) ( error)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map":{
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world, each subsequent call to this command will display the next 20 locations",
			callback:    commandMapf,
		},
		"mapb":{
			name: "mapb",
			description: "displays the previous 20 locations, if first page it will give an error",
			callback: commandMapB,
		},
		"explore":{
			name: "explore",
			description: "explore the pokemon in that location-area",
			callback: commandExplore,
		},
		"catch":{
			name: "catch",
			description: "catch the pokemon randomly",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect",
			description: "inspect the pokemon",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "display the pokedex",
			callback: commandPokedex,
		},
	}
}