package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
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

var locationAreaUrl = "https://pokeapi.co/api/v2/location-area/"

func GetLocationAreaPage(url string) (PaginatedLocationArea, error) {
	plocarea := PaginatedLocationArea{}

	if url == "" {
		url = locationAreaUrl + "?offset=0&limit=20"
	}

	data, err := pokeGet(url)
	if err != nil {
		return plocarea, fmt.Errorf("GET error: %w", err)
	}

	if err := json.Unmarshal(data, &plocarea); err != nil {
		return plocarea, fmt.Errorf("Error unmarshalling data: %w", err)
	}

	return plocarea, nil
}

func GetLocationArea(name string) (LocationArea, error) {
	locarea := LocationArea{}

	if name == "" {
		return locarea, errors.New("No id or name specified")
	}

	data, err := pokeGet(locationAreaUrl + name)
	if err != nil {
		return locarea, fmt.Errorf("GET error: %w", err)
	}

	if err := json.Unmarshal(data, &locarea); err != nil {
		return locarea, fmt.Errorf("Error unmarshalling data: %w", err)
	}

	return locarea, nil
}
