package main

import (
	"github.com/Levi-1103/pokecli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	pokeapi.DisplayPokemon(args[0])
	return nil
}
