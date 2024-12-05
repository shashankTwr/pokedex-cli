package main

import (
	"github.com/shashankTwr/pokedexcli/internal/pokeapi"
)


func commandMap() error {
    return pokeapi.GetNextLocations()
}

func commandMapB() error {
    return pokeapi.GetPreviousLocations()
}

func main(){
	startRepl()
}