package main

import (
	"time"

	"github.com/Levi-1103/pokecli/internal/pokecache"
)

type config struct {
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	cache                   pokecache.Cache
}

func main() {

	cfg := config{
		nextLocationAreaURL:     nil,
		previousLocationAreaURL: nil,
		cache:                   pokecache.NewCache(500 * time.Millisecond),
	}

	startRepl(&cfg)

}
