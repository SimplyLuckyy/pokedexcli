package main

import (
	"fmt"
	"os"
	"errors"
	"math/rand"
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
	
	area, err := cfg.pokeapiClient.ListLocaINFO(args[0])
	if err != nil {return err}
	
	encounters := area.PokemonEncounters
	
	for _, encounter := range encounters {fmt.Println("- ", encounter.Pokemon.Name)}
	
	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Missing argument: catch <pokemon>")
	}
	poke, err := cfg.pokeapiClient.ListPokeINFO(args[0])
	if err != nil {return err}

	fmt.Printf("Throwing a Pokeball at %s...\n", poke.Name)
	catchRate := rand.Intn(poke.BaseExperience)

	if catchRate > 35 {
		fmt.Printf("%s escaped!\n", poke.Name)
		return nil
	}
	
	fmt.Printf("%s was caught!\n", poke.Name)
	fmt.Println("You may now inspect it with the inspect command")
	cfg.pokedex[args[0]] = poke
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Missing argument: inspect <pokemon>")
	}
	pokeName := args[0]

	pokeData, ok := cfg.pokedex[pokeName]
	if !ok {return errors.New("Pokemon not in Pokedex")}

	fmt.Printf("Name: %s\n", pokeData.Name)
	fmt.Printf("Height: %v\n", pokeData.Height)
	fmt.Printf("Weight: %v\n", pokeData.Weight)
	fmt.Println("Stats:")
	for _, statINFO := range pokeData.Stats {
		fmt.Printf("  - %s: %v\n", statINFO.Stat.Name, statINFO.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeINFO := range pokeData.Types {
		fmt.Printf("  - %s\n", typeINFO.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) == 0 {return errors.New("You have not caught any pokemon")}

	fmt.Println("Your Pokedex:")
	for poke := range cfg.pokedex {
		fmt.Printf("  - %s\n", poke)
	}
	return nil
}