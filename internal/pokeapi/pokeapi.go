package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/espenronnevik/bootdev-pokedex/internal/pokecache"
)

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Second)
}

func pokeGet(url string) ([]byte, error) {
	var data []byte

	data, cached := cache.Get(url)
	if !cached {
		res, err := http.Get(url)
		if err != nil {
			return data, fmt.Errorf("Network error: %w", err)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return data, fmt.Errorf("Error reading data: %w", err)
		}
		cache.Add(url, data)
	}
	return data, nil
}
