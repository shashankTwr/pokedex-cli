package main

import (
	"bufio"
	"fmt"
	"os"

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

