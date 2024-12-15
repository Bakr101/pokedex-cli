package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error{
	if len(args) != 1{
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	_, ok := c.pokedex.GetPokemon(pokemonName)
	if ok {
		return fmt.Errorf("%v is already caught", pokemonName)
	}
	pokemon, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	randomNumber := rand.Intn(pokemon.BaseExperience)
	halfExperience := pokemon.BaseExperience / 2
	var catchStatus string
	if randomNumber >= halfExperience{
		catchStatus = pokemon.Name + " was caught!"
		c.pokedex.AddPokemon(pokemon)
	}else{
		catchStatus = pokemon.Name + " escaped!"
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	fmt.Printf("%v\n", catchStatus)
	fmt.Println("You may now inspect it with the inpect command.")

	return nil
}

