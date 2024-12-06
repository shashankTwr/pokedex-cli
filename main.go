package main

import (
	"time"

	"github.com/shashankTwr/pokedexcli/internal/pokeapi"
)

func main(){

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	// timeIndia, _ := time.LoadLocation("Asia/Kolkata")
	// fmt.Println(timeIndia)
	startRepl(cfg)
}