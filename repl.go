package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/keyinora/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		area_name := words[1]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, area_name)
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
	if text == "" {
		return []string{}
	}
	message := strings.ToLower(text)
	// only returns letters and spaces, removes everything else.
	message = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, message)
	return strings.Fields(message)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
