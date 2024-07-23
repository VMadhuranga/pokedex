package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/VMadhuranga/pokedex/pokecache"
)

type locationAreas struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
	}
}

type locationArea struct {
	Name              string
	PokemonEncounters []struct {
		Pokemon struct {
			Name string
		}
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string
		}
	}
	Types []struct {
		Type struct {
			Name string
		}
	}
}

var pCache = pokecache.NewCache(5 * time.Second)

func GetLocationAreas(url string) locationAreas {
	areas := locationAreas{}
	entry, ok := pCache.GetEntry(url)
	if ok {
		fmt.Println("(cached results)")
		if err := json.Unmarshal(entry, &areas); err != nil {
			log.Fatal(err)
		}
		return areas
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	pCache.AddEntry(url, body)
	if err := json.Unmarshal(body, &areas); err != nil {
		log.Fatal(err)
	}
	return areas
}

func GetLocationArea(url string) (locationArea, error) {
	area := locationArea{}
	entry, ok := pCache.GetEntry(url)
	if ok {
		fmt.Println("(cached results)")
		if err := json.Unmarshal(entry, &area); err != nil {
			log.Fatal(err)
		}
		return area, nil
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode == 404 {
		return locationArea{}, errors.New("area not found")
	}
	if res.StatusCode > 299 {
		log.Fatalf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	pCache.AddEntry(url, body)
	if err := json.Unmarshal(body, &area); err != nil {
		log.Fatal(err)
	}
	return area, nil
}

func GetPokemon(url string) (Pokemon, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode == 404 {
		return Pokemon{}, errors.New("pokemon not found")
	}
	if res.StatusCode > 299 {
		log.Fatalf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	p := Pokemon{}
	if err := json.Unmarshal(body, &p); err != nil {
		log.Fatal(err)
	}
	return p, nil
}
