package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PaginatedLocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var LocationAreaUrl = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

func GetLocationArea(url string) (PaginatedLocationArea, error) {
	var data []byte

	locarea := PaginatedLocationArea{}
	data, cached := cache.Get(url)

	if !cached {
		res, err := http.Get(url)
		if err != nil {
			return locarea, fmt.Errorf("Network error: %w", err)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return locarea, fmt.Errorf("Error reading data: %w", err)
		}
		cache.Add(url, data)
	}

	if err := json.Unmarshal(data, &locarea); err != nil {
		return locarea, fmt.Errorf("Error unmarshalling data: %w", err)
	}

	return locarea, nil
}
