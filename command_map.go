package main

import "github.com/Levi-1103/pokecli/internal/pokeapi"

func commandMap(cfg *config) error {
	next, prev, err := pokeapi.DisplayLocations(cfg.nextLocationAreaURL)

	cfg.nextLocationAreaURL = next
	cfg.previousLocationAreaURL = prev

	return err
}
