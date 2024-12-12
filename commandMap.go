package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config) error{
	locations, err := c.pokeapiClient.ListLocations(c.nextURL)
	if err != nil {
		return err
	}
	for _, location := range locations.Results{
		fmt.Printf("%v\n", location.Name)
	}
	c.nextURL = locations.Next
	c.prevURL = locations.Previous

	return nil
}

func commandMapb(c *config) error{
	if c.prevURL == nil {
		return errors.New("you're on the first page.")
	}
	locations, err := c.pokeapiClient.ListLocations(c.prevURL)
	if err != nil {
		return err
	}
	for _, location := range locations.Results{
		fmt.Printf("%v\n", location.Name)
	}
	c.nextURL = locations.Next
	c.prevURL = locations.Previous

	return nil
}