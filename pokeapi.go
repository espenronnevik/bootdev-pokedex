package main

import (
	"encoding/json"
	"fmt"
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

var urlLocationArea = "https://pokeapi.co/api/v2/location-area/"

func getLocationArea(url string) (PaginatedLocationArea, error) {
	locarea := PaginatedLocationArea{}

	res, err := http.Get(url)
	if err != nil {
		return locarea, fmt.Errorf("Network error: %w", err)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locarea); err != nil {
		return locarea, fmt.Errorf("JSON decoding error: %w", err)
	}

	return locarea, nil
}
