package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: catch <pokemon-name>")
	}

	p_name := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", p_name)
	url = normalizeURL(url)

	var data []byte
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("(cached result)")
		data = cachedData
	} else {
		fmt.Println("(cache miss — fetching from API...)")
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error fetching Pokemon data from API: %s", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error reading the response's body: %s", err)
		}
		data = body

		cfg.cache.Add(url, data)
	}

	var pkm Pokemon
	if err := json.Unmarshal(data, &pkm); err != nil {
		return fmt.Errorf("Error parsing JSON: %s", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pkm.Name)

	rand.Seed(time.Now().UnixNano())
	catchChance := 100 - pkm.baseExp/10
	if catchChance < 10 {
		catchChance = 10
	}
	if rand.Intn(100) < catchChance {
		fmt.Printf("%s was caught!\n", pkm.Name)
		if cfg.pokedex == nil {
			cfg.pokedex = make(map[string]Pokemon)
		}
		cfg.pokedex[pkm.Name] = pkm
	} else {
		fmt.Printf("%s escaped!\n", pkm.Name)
	}

	return nil
}