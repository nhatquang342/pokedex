package main

import (
	"time"
	"github.com/nhatquang342/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)

    cfg := &config{
        cache: cache,
    }

    startRepl(cfg)
}