package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jifpbj/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", commandName, err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show the help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "map out 20 areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map back",
			description: "map previous 20 areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore an area for pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch the pokemon",
			callback:    commandCatch,
		},
	}
}
