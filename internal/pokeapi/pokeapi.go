package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func GetLocationArea(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	payload := LocationArea{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return LocationArea{}, err
	}

	return payload, nil
}

type PokemonInLocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetPokemonInLocationArea(url string) (PokemonInLocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokemonInLocationArea{}, err
	}
	defer res.Body.Close()

	payload := PokemonInLocationArea{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return PokemonInLocationArea{}, err
	}

	return payload, nil
}

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetPokemon(url string) (Pokemon, error) {
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	payload := Pokemon{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return Pokemon{}, err
	}

	return payload, nil
}
