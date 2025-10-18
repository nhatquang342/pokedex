package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/nhatquang342/pokedex/internal/pokecache"
)

func startRepl(cfg *config) {
	
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned := reader.Scan()
		if !scanned {
			break
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex map[string]Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the location areas",
			callback:    commandMap,
		},
		"mapb": {
            name:        "mapb",
            description: "Show the previous 20 location areas",
            callback:    commandMapb,
        },
		"explore": {
			name:        "explore",
            description: "Show pokemons of an area",
            callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
            description: "Catch a pokemon",
            callback:    commandCatch,
		},
	}
}
