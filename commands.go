package main

import (
	"fmt"
	"os"
	"errors"
)

const (baseURL = "https://pokeapi.co/api/v2/location-area/canalave-city-area")

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range checkCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config, args []string) error {
	loca, err := cfg.pokeapiClient.ListLoca(cfg.next)
	if err != nil {return err}

	cfg.next = loca.Next
	cfg.prev = loca.Prev

	for _, l := range loca.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prev == nil {return errors.New("You're on the first page")}

	loca, err := cfg.pokeapiClient.ListLoca(cfg.prev)
	if err != nil {return err}

	cfg.next = loca.Next
	cfg.prev = loca.Prev

	for _, l := range loca.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandExplore(cfg *config, args []string) error {
	
	if len(args) != 1 {
		return errors.New("Missing argument: explore <area-name>")
	}
	
	poke, err := cfg.pokeapiClient.ListLocaINFO(args[0])
	if err != nil {return err}
	
	encounters := poke.PokemonEncounters
	
	for _, encounter := range encounters {fmt.Println("- ", encounter.Pokemon.Name)}
	
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

