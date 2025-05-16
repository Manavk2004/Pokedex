package main

import(
	"time"
	"github.com/Manavk2004/Pokedex/internal/pokeapi"
)

func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}