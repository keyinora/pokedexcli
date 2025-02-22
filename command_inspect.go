package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name\n")
	}

	name := args[0]
	data, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("You do not have this pokemon in your pokedex yet\n")
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %v\n", data.Height)
	fmt.Printf("Weight: %v\n", data.Weight)
	fmt.Print("Stats:\n")
	for _, s := range data.Stats {
		fmt.Printf(" -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, t := range data.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}
