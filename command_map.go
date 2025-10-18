package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreasResp struct {
    Count    int `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config, args []string) error {
	url := cfg.next
    if url == "" {
        url = "https://pokeapi.co/api/v2/location-area/"
    }
	
	url = normalizeURL(url)
	if data, found := cfg.cache.Get(url); found {
		fmt.Println("(cached result)")
		return printLocationAreas(cfg, data)
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

	return printLocationAreas(cfg, data)
}

func commandMapb(cfg *config, args []string) error {
	url := cfg.previous
    if url == "" {
        fmt.Println("You're on the first page.")
        return nil
    }
	
	url = normalizeURL(url)
	if data, found := cfg.cache.Get(url); found {
		fmt.Println("(cached result)")
		return printLocationAreas(cfg, data)
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
	
	return printLocationAreas(cfg, data)
}

func printLocationAreas(cfg *config, data []byte) error {
	var resp locationAreasResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.next = resp.Next
	cfg.previous = resp.Previous
	return nil
}