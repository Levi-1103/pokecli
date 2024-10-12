package main

import (
	"time"

	"github.com/Levi-1103/pokecli/internal/pokecache"
	"github.com/Levi-1103/pokecli/internal/pokedex"
)

type config struct {
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	cache                   pokecache.Cache
	pokedex                 pokedex.Pokedex
}

func main() {

	cfg := config{
		nextLocationAreaURL:     nil,
		previousLocationAreaURL: nil,
		cache:                   pokecache.NewCache(500 * time.Millisecond),
		pokedex:                 pokedex.NewPokedex(),
	}

	startRepl(&cfg)

}
