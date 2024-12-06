package main

import (
	"time"

	"github.com/shashankTwr/pokedexcli/internal/pokeapi"
	"github.com/shashankTwr/pokedexcli/internal/pokecache"
)
const baseURL = "https://pokeapi.co/api/v2/location-area/"


var Cfg *pokeapi.Config = &pokeapi.Config{
	Count: 0,
    Next: "",
    Previous: "",
    Results: nil,
}

const interval = 5 * time.Second
var Cache *pokecache.Cache = pokecache.NewCache(interval)
func main(){
	startRepl()
}