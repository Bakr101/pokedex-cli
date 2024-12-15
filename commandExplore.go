package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error{
	if len(args) != 1{
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]
	locationArea, err := c.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemons in %v:\n", locationArea.Name)
	for _, pokeStruct := range locationArea.PokemonEncounters{
		fmt.Printf("%v\n", pokeStruct.Pokemon.Name)
	}


	return nil
}

