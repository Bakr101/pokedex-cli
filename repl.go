package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Bakr101/pokedex-cli-v1/internal/pokeapi"
)

type cliCommand struct{
	name		string
	description	string
	callback	func(*config, ...string) error
}

type config struct{
	pokeapiClient	pokeapi.Client
	pokedex			Pokedex
	nextURL			*string
	prevURL			*string
}



func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"map":{
			name: "map",
			description: "Get the next page of locations",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapb,
		},
		"explore":{
			name: "explore {location-name}",
			description: "Explore which pokemons are available in an area",
			callback: commandExplore,
		},
		"catch":{
			name: "catch {pokemon-name}",
			description: "Tries to catch a pokemon and add it to the Pokedex",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect {pokemon-name}",
			description: "Lists all the details about a pokemon you caught",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "Lists all the pokemons in your Pokedex",
			callback: commandPokedex,
		},
	}
}

func startRepl(c *config){
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		text := cleanInput(reader.Text())
		if len(text) == 0{
			continue
		}
		commandName := text[0]
		args := []string{}
		if len(text) > 1{
			args = text[1:] 
		}
		
		if command, ok := commands[commandName]; ok{
			err := command.callback(c, args...)
			if err != nil{
				fmt.Println(err)
			}
		}else{
			fmt.Println("--> Unkown command, Please try help command to see all available commands.")
			continue
		}
		
		
	}
}

func cleanInput(text string) []string{
	textLowerCase := strings.ToLower(text)
	textSlice := strings.Fields(textLowerCase)
	return textSlice
}