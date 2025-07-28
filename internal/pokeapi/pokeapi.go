package pokeapi

import (
	"time"

	"github.com/espenronnevik/bootdev-pokedex/internal/pokecache"
)

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Second)
}
