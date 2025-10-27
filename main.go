package main

import (

	"github.com/simplyluckyy/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

