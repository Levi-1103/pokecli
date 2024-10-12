package main

import "github.com/Levi-1103/pokecli/internal/pokeapi"

func commandPreviousMap(cfg *config) error {

	next, prev, err := pokeapi.DisplayLocations(cfg.previousLocationAreaURL, cfg.cache)

	cfg.nextLocationAreaURL = next
	cfg.previousLocationAreaURL = prev

	return err
}
