package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := cliConfig{
		initial:  "https://pokeapi.co/api/v2/location-area/",
		next:     "",
		previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		// waits for user to type and press enter
		scanner.Scan()
		// gets the text that was typed
		res := scanner.Text()
		//split the words
		words := cleanInput(res)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if command, exists := getCommands()[commandName]; exists {
			err := command.callback(config)
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

/*registry of commands*/
type cliCommand struct {
	name        string
	description string
	callback    func(cliConfig) error
}

/*congif struct*/
type cliConfig struct {
	initial  string
	next     string
	previous string
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
			description: "Display 20 locations in the Pokemon world. Each call with get the next 20.",
			callback:    commandMap,
		},
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
