package main

import (
	"fmt"
	"os"
	"errors"
)



func commandHelp(*cfg config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range checkCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(*cfg config) error {
	
}

func commandMapb(*cfg config) error {

}

func commandExit(*cfg config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

