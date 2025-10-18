package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your pokedex is empty. You haven't caught any pokemons")
		return nil
	}
	
	names := make([]string, 0, len(cfg.pokedex))
	for name := range cfg.pokedex {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	for _, name := range names {
		fmt.Println("-", name)
	}
	return nil
}