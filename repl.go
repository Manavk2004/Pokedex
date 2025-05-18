package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/Manavk2004/Pokedex/internal/pokeapi"
)


type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config){
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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
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
	callback func(cfg *config, args ...string) error
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
		"map": {
			name: "map", 
			description: "Get the next page of locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations", 
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Explore the locations/pokemon",
			callback: commandExplore,
		},
	}
}

func cleanInput(text string) []string{
	lowerWord := strings.ToLower(text)
	splitWord := strings.Fields(lowerWord)
	return splitWord
}
