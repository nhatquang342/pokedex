package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore (cfg *config, args []string) error {
	if len(args) <= 0 {
		fmt.Printf("Usage: explore <location-area-name>")
	}
	name := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", name)

	url = normalizeURL(url)
	if data, found := cfg.cache.Get(url); found {
		fmt.Println("(cached result)")
		return printPokemon(data)
	}

	fmt.Println("(cache miss — fetching from API...)")
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error making HTTP request: %s", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading the response's body: %s", err)
	}

	cfg.cache.Add(url, data)
	return printPokemon(data)
}

func printPokemon (data []byte) error {
	var resp locationArea
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	for _, pkm := range resp.PokemonEncounters {
		fmt.Println(pkm.Pokemon.Name)
	}

	return nil
}