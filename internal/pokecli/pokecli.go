package pokecli

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

type CliCommand struct {
	Name, Description string
}

type CliConfig struct {
	RootLocationAreaUrl, nextLocationAreaUrl, prevLocationAreaUrl, RootPokemonUrl string
	CliCommands                                                                   []CliCommand
	Cache                                                                         *pokecache.Cache
	Pokedex                                                                       map[string]pokeapi.Pokemon
}

func (c *CliConfig) InitExitCmd() error {
	defer os.Exit(0)
	return nil
}

func (c *CliConfig) InitHelpCmd() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range c.CliCommands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}

func (c *CliConfig) InitMapCmd() error {
	url := c.nextLocationAreaUrl
	if url == "" {
		url = c.RootLocationAreaUrl
	}

	lArea := pokeapi.LocationArea{}

	cachedData, ok := c.Cache.Get(url)
	if !ok {
		locationArea, err := pokeapi.GetLocationArea(url)
		if err != nil {
			return err
		}

		marshaledData, err := json.Marshal(locationArea)
		if err != nil {
			return err
		}

		c.Cache.Add(url, marshaledData)
		cachedData = marshaledData
	}

	err := json.Unmarshal(cachedData, &lArea)
	if err != nil {
		return err
	}

	c.nextLocationAreaUrl = lArea.Next
	c.prevLocationAreaUrl = lArea.Previous

	if c.nextLocationAreaUrl == "" {
		c.nextLocationAreaUrl = c.prevLocationAreaUrl
	}

	for _, v := range lArea.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	return nil
}

func (c *CliConfig) InitMapbCmd() error {
	url := c.prevLocationAreaUrl
	if url == "" {
		url = c.RootLocationAreaUrl
	}

	lArea := pokeapi.LocationArea{}

	cachedData, ok := c.Cache.Get(url)
	if !ok {
		locationArea, err := pokeapi.GetLocationArea(url)
		if err != nil {
			return err
		}

		marshaledData, err := json.Marshal(locationArea)
		if err != nil {
			return err
		}

		c.Cache.Add(url, marshaledData)
		cachedData = marshaledData
	}

	err := json.Unmarshal(cachedData, &lArea)
	if err != nil {
		return err
	}

	c.nextLocationAreaUrl = lArea.Next
	c.prevLocationAreaUrl = lArea.Previous

	if c.prevLocationAreaUrl == "" {
		c.prevLocationAreaUrl = c.RootLocationAreaUrl
	}

	for _, v := range lArea.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	return nil
}

func (c *CliConfig) InitExploreCmd(locationArea string) error {
	pokeInLArea := pokeapi.PokemonInLocationArea{}

	cachedData, ok := c.Cache.Get(locationArea)
	if !ok {
		pokemonInLocationArea, err := pokeapi.GetPokemonInLocationArea(c.RootLocationAreaUrl + locationArea)
		if err != nil {
			return err
		}

		marshaledData, err := json.Marshal(pokemonInLocationArea)
		if err != nil {
			return err
		}

		c.Cache.Add(locationArea, marshaledData)
		cachedData = marshaledData
	}

	err := json.Unmarshal(cachedData, &pokeInLArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\nFound Pokemon:\n", locationArea)
	for _, encounter := range pokeInLArea.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}

func (c *CliConfig) InitCatchCmd(pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	random := rand.Intn(3)
	if random == 1 {
		fmt.Printf("%v was caught!\n", pokemonName)

		pokemon, err := pokeapi.GetPokemon(c.RootPokemonUrl + pokemonName)
		if err != nil {
			return err
		}

		c.Pokedex[pokemonName] = pokemon

		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%v escaped!!\n", pokemonName)
	}

	return nil
}

func (c *CliConfig) InitInspectCmd(pokemonName string) error {
	pokemon, ok := c.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name: ", pokemon.Name)
		fmt.Println("Height: ", pokemon.Height)
		fmt.Println("Weight: ", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, s := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types: ")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}

	}

	return nil
}

func (c *CliConfig) InitPokedexCmd() error {
	fmt.Println("Your Pokedex:")
	for _, poke := range c.Pokedex {
		fmt.Printf("  - %v\n", poke.Name)
	}

	return nil
}
