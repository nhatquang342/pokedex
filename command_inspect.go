package main

import (
	//"encoding/json"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: inspect <pokemon-name>")
	}

	p_name := args[0]
	/*
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", p_name)
	url = normalizeURL(url)

	if cachedData, found := cfg.cache.Get(url); !found {
		return fmt.Errorf("you have not caught that pokemon")
	} else {
		var pkm Pokemon
		if err := json.Unmarshal(cachedData, &pkm); err != nil {
			return fmt.Errorf("Error parsing JSON: %s", err)
		}
		
		fmt.Println("Name:", pkm.Name)
		fmt.Println("Height:", pkm.Height)
		fmt.Println("Weight:", pkm.Weight)
	}
	*/

	if pkm, ok := cfg.pokedex[p_name]; ok {
		fmt.Println("Name:", pkm.Name)
		fmt.Println("Height:", pkm.Height)
		fmt.Println("Weight:", pkm.Weight)
		return nil
	} else {
		return fmt.Errorf("you have not caught that pokemon")
	}
}