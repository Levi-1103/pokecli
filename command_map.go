package main

import "github.com/Levi-1103/pokecli/internal/pokeapi"

func commandMap(cfg *config, args ...string) error {
	next, prev, err := pokeapi.DisplayLocations(cfg.nextLocationAreaURL, cfg.cache)

	cfg.nextLocationAreaURL = next
	cfg.previousLocationAreaURL = prev

	return err
}
