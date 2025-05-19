package main

import(
	"fmt"
	"errors"
)


func commandPokeDex (cfg *config, args ...string) error{
	if len(args) >= 1{
		return errors.New("Not a valid command")
	}
	for _, p := range cfg.caughtPokemon{
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
