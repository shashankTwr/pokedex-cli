package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)


func commandHelp(cfg *config, args ...string) error {
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

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}


func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationsResp.Next
	cfg.prevLocationsURL = &locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationResp.Next
	cfg.prevLocationsURL = &locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationsResp, err := cfg.pokeapiClient.ExploreLocations(&args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationsResp.Location.Name)
	if len(locationsResp.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found")
		return nil
	}
	fmt.Printf("Found %d pokemon:\n", len(locationsResp.PokemonEncounters))
	for _, pokemon := range locationsResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(&args[0])
	if err != nil {
		return err
	}
	baseExp := pokemonResp.BaseExperience
	chance := rand.Intn(10)
	if chance == (baseExp % 10) {
		pokedex.addCaught(args[0], pokemonResp)
		fmt.Printf("You caught %s!\n", pokemonResp.Name)
	}else{
		fmt.Printf("You failed to catch %s\n", pokemonResp.Name)
	}
	// fmt.Printf("%v\n",baseExp)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon, exists := pokedex.caught[args[0]]
	if !exists {
		return errors.New("pokemon not caught")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t%s: %d\n", stat.Stat.Name,stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	
	if pokedex.caught == nil {
		fmt.Println("No pokemon caught yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name, pokemon := range pokedex.caught {
		fmt.Printf("%s: %s\n", name, pokemon.Name)
	}
	return nil
}