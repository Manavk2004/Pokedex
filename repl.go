package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)



func startRepl(){
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	
	for{
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0{
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil{
				fmt.Println(err)
			}
			continue
		}else{
			fmt.Println("Unknown command")
			continue
		}
	}
}


type cliCommand struct{
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit", 
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}

func cleanInput(text string) []string{
	lowerWord := strings.ToLower(text)
	splitWord := strings.Fields(lowerWord)
	return splitWord
}
