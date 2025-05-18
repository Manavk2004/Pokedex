package main

import(
	"errors"
	"fmt"
)


func commandExplore(cfg *config, args ...string) error{
	if len(args) != 1 {
		return errors.New("You must provide a location")
	}
	locationName := args[0]
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil{
		return err
	}
	fmt.Printf("Exploring %s...", location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters{
		fmt.Printf(enc.Pokemon.Name)
	}
	return nil


}