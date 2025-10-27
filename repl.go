package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"

	"github.com/simplyluckyy/pokedexcli/pokeapi"
)

type cliCommand struct {
	name		string
	description string
	callback	func(*config) error
}

type config struct {
	pokeapiClient	pokeapi.Client
	next			*string
	prev 	 		*string
}

func checkCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback:	 commandHelp,
		},
		"map": {
			name:		 "map",
			description: "Display the next twenty locations",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Display the previous twenty locations",
			callback:	 commandMapb,
		},
		"exit": {
			name:		 "exit",
			description: "Exit the pokedex",
			callback:	 commandExit,
		},
	}
}


func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {continue}
		command, ok := checkCommands()[input[0]]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	fixedText := strings.Fields((strings.ToLower(text)))
	return fixedText
}

