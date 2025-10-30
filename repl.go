package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"

	"github.com/simplyluckyy/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name		string
	description string
	callback	func(*config, []string) error
}

type config struct {
	pokeapiClient	pokeapi.Client
	next			*string
	prev 	 		*string
	pokedex			map[string]pokeapi.PokeAPIPokeINFO
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
		"explore": {
			name:		 "explore",
			description: "Explore a location",
			callback:	 commandExplore,
		},
		"catch": {
			name:		 "catch",
			description: "Catch a pokemon",
			callback:	 commandCatch,
		},
		"inspect": {
			name:		 "inspect",
			description: "Inspect a caught pokemon",
			callback:	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "Vew your caught pokemon",
			callback:	 commandPokedex,
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
		command, args  := cleanInput(scanner.Text())
		if len(command) == 0 {continue}
		com, ok := checkCommands()[command]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		} else {
			err := com.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) (command string, args []string) {
	fixedText := strings.Fields((strings.ToLower(text)))
	command = fixedText[0]
	args = fixedText[1:]
	return command, args
}

