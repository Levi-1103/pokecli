package main

import "github.com/Levi-1103/pokecli/internal/pokedex"

func commandPokedex(cfg *config, args ...string) error {

	err := pokedex.ShowPokedex(cfg.pokedex)

	return err
}
