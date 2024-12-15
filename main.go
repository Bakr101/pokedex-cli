package main

import (
	"time"

	"github.com/Bakr101/pokedex-cli-v1/internal/pokeapi"
)



func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	pokedexStore := NewPokedex()
	config := &config{
		pokeapiClient: pokeClient,
		pokedex: pokedexStore,
	}
	startRepl(config)
}


