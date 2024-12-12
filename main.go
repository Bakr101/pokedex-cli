package main

import (
	"time"

	"github.com/Bakr101/pokedex-cli-v1/internal/pokeapi"
)



func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
		
	}
	startRepl(config)
	
}


