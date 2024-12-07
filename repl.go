package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shashankTwr/pokedexcli/internal/pokeapi"
)

type config struct{
	pokeapiClient pokeapi.Client	
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		var argsName string
		commandName := words[0]
		if len(words) == 2 {
			argsName = words[1]
		}
		

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, argsName)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

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
	}
}