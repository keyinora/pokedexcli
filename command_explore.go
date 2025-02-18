package main

import "fmt"

func commandExplore(cfg *config, area_name string) error {
	println()
	fmt.Printf("Exploring %s...", area_name)
	println()
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(area_name)
	if err != nil {
		return err
	}
	for _, data := range exploreResp.PokemonEncounters {
		fmt.Printf("-%s", data.Pokemon.Name)
		println()
	}
	return nil
}
