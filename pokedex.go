package main

import (
	"sync"

	"github.com/Bakr101/pokedex-cli-v1/internal/pokeapi"
)

type Pokedex struct{
	pokemons		map[string]pokeapi.Pokemon
	mu				*sync.Mutex
}

func NewPokedex()Pokedex{
	return Pokedex{
		pokemons: make(map[string]pokeapi.Pokemon),
		mu: &sync.Mutex{},
	}
}

func (p *Pokedex) AddPokemon(pokemon pokeapi.Pokemon){
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pokemons[pokemon.Name] = pokemon
}

func (p *Pokedex) GetPokemon(pokemonName string) (pokeapi.Pokemon, bool){
	p.mu.Lock()
	defer p.mu.Unlock()
	pokemon, ok := p.pokemons[pokemonName]
	if !ok{
		return pokeapi.Pokemon{}, ok
	}
	return pokemon, ok
}

// Add Feature to get all pokemons in pokedex