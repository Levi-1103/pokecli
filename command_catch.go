package main

import "github.com/Levi-1103/pokecli/internal/pokedex"

func commandCatch(cfg *config, args ...string) error {
	err := pokedex.CatchPokemon(args[0], cfg.pokedex)
	return err
}
