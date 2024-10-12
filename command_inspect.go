package main

import "github.com/Levi-1103/pokecli/internal/pokedex"

func commandInspect(cfg *config, args ...string) error {
	err := pokedex.InspectPokemon(args[0], cfg.pokedex)
	return err
}
