package main

import(
	"time"
	"github.com/Manavk2004/Pokedex/internal/pokeapi"
)

func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)

}