package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error{
	if len(args) != 1{
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := c.pokedex.GetPokemon(pokemonName)
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	
	
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats{
		fmt.Printf("	-%v: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _,typ := range pokemon.Types{
		fmt.Printf("	-%v\n", typ.Type.Name)
	}
	
	return nil
}

