package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error{
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokedex.pokemons{
		fmt.Printf(" -%v\n", pokemon.Name)
	}
	
	return nil
}