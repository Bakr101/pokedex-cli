package main

import "fmt"

func commandHelp(c *config, args ...string)error{
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, command := range commands{
		fmt.Printf("%s : %s\n" ,command.name, command.description)
	}
	fmt.Println()
	return nil
}