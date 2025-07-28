package pokeapi

import (
	"encoding/json"
	"errors"
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

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var LocationAreaUrl = "https://pokeapi.co/api/v2/location-area/"

func GetLocationAreaPage(url string) (PaginatedLocationArea, error) {
	var data []byte

	if url == "" {
		url = LocationAreaUrl + "?offset=0&limit=20"
	}

	plocarea := PaginatedLocationArea{}
	data, cached := cache.Get(url)

	if !cached {
		res, err := http.Get(url)
		if err != nil {
			return plocarea, fmt.Errorf("Network error: %w", err)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return plocarea, fmt.Errorf("Error reading data: %w", err)
		}
		cache.Add(url, data)
	}

	if err := json.Unmarshal(data, &plocarea); err != nil {
		return plocarea, fmt.Errorf("Error unmarshalling data: %w", err)
	}

	return plocarea, nil
}

func GetLocationArea(name string) (LocationArea, error) {
	var data []byte

	locarea := LocationArea{}
	if name == "" {
		return locarea, errors.New("No id or name specified")
	}

	url := LocationAreaUrl + name

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
