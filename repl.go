package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct{
	name		string
	description	string
	callback	func() error
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
	}
}

func startRepl(){
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		text := cleanInput(reader.Text())
		if len(text) == 0{
			continue
		}
		input := text[0]
		if command, ok := commands[input]; ok{
			err := command.callback()
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