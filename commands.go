package main

import (
	"fmt"
	"os"
	"errors"
)



func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range checkCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	locaResp, err := cfg.pokeapiClient.ListLoca(cfg.next)
	if err != nil {return err}

	cfg.next = locaResp.Next
	cfg.prev = locaResp.Prev

	for _, loca := range locaResp.Results {
		fmt.Println(loca.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prev == nil {return errors.New("You're on the first page")}

	locaResp, err := cfg.pokeapiClient.ListLoca(cfg.prev)
	if err != nil {return err}

	cfg.next = locaResp.Next
	cfg.prev = locaResp.Prev

	for _, loca := range locaResp.Results {
		fmt.Println(loca.Name)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

