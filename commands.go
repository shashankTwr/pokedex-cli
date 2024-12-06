package main

import (
	"fmt"
	"os"

	"github.com/shashankTwr/pokedexcli/internal/pokeapi"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}


func commandMap() error {
	fmt.Println(Cfg.Previous)
    return pokeapi.GetNextLocations(Cache, Cfg,baseURL)
}

func commandMapB() error {
    fmt.Println(Cfg.Previous)
	return pokeapi.GetPreviousLocations(Cache,Cfg)

}