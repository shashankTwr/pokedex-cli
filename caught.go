package main

import "github.com/shashankTwr/pokedexcli/internal/pokeapi"

type Pokedex struct {
	caught map[string]pokeapi.RespPokemon
}
var pokedex *Pokedex = &Pokedex{}
func (p *Pokedex) addCaught(name string, pokemon pokeapi.RespPokemon) {
	if p.caught == nil {
		p.caught = make(map[string]pokeapi.RespPokemon)
	}
	p.caught[name] = pokemon
}