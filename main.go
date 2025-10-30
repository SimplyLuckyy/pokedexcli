package main

import (
	"time"

	"github.com/simplyluckyy/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: map[string]pokeapi.PokeAPIPokeINFO{},
	}

	startRepl(cfg)
}

