package pokecli

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"

	"github.com/VMadhuranga/pokedex/pokeapi"
)

type cliCommand struct {
	name        string
	description string
}

func fmtCmdDescription(cmd, description string) string {
	return fmt.Sprintf("%v%v%v", cmd, strings.Repeat(" ", 20-len(cmd)), description)
}

func InitPokeCliMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: fmtCmdDescription("help", "displays a help message"),
		},
		"exit": {
			name:        "exit",
			description: fmtCmdDescription("exit", "exit the pokedex"),
		},
		"map": {
			name:        "map",
			description: fmtCmdDescription("map", "displays the names of next 20 location areas"),
		},
		"mapb": {
			name:        "mapb",
			description: fmtCmdDescription("mapb", "displays the names of previous 20 location areas"),
		},
		"explore": {
			name:        "explore",
			description: fmtCmdDescription("explore [area]", "displays list of all the pokemon in a given area"),
		},
		"catch": {
			name:        "catch",
			description: fmtCmdDescription("catch [pokemon]", "catch Pokemon and add them to the pokedex"),
		},
		"inspect": {
			name:        "inspect",
			description: fmtCmdDescription("inspect [pokemon]", "displays details about pokemon"),
		},
		"pokedex": {
			name:        "pokedex",
			description: fmtCmdDescription("pokedex", "displays list of captured pokemon"),
		},
	}
}

func Prompt() {
	fmt.Print("pokedex > ")
}

func RunCommand(cliCmdMap map[string]cliCommand, cmd cliCommand, arg string) {
	switch cmd.name {
	case "help":
		runCmdHelp(cliCmdMap)
	case "exit":
		runCmdExit()
	case "map":
		runCmdMap()
	case "mapb":
		runCmdMapB()
	case "explore":
		runCmdExplore(arg)
	case "catch":
		runCmdCatch(arg)
	case "inspect":
		runCmdInspect(arg)
	case "pokedex":
		runCmdPokedex()
	}
}

func runCmdHelp(cliCmdMap map[string]cliCommand) {
	fmt.Print("welcome to the pokedex!\n")
	fmt.Print("\nusage:\n\n")
	keys := []string{}
	for key := range cliCmdMap {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	for _, k := range keys {
		cmd := cliCmdMap[k]
		fmt.Println(cmd.description)
	}
	fmt.Println()
}

func runCmdExit() {
	os.Exit(0)
}

const locationAreaUrl = "https://pokeapi.co/api/v2/location-area/"

type mapConfig struct {
	nextUrl string
	prevUrl string
}

func (mc *mapConfig) updateMapConfig(next, prev string) {
	mc.nextUrl = next
	mc.prevUrl = prev
}

var mc = mapConfig{
	nextUrl: fmt.Sprintf("%v?offset=0&limit=20", locationAreaUrl),
}

func runCmdMap() {
	if len(mc.nextUrl) == 0 {
		fmt.Println("you are on the last page")
	} else {
		areas := pokeapi.GetLocationAreas(mc.nextUrl)
		mc.updateMapConfig(areas.Next, areas.Previous)
		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}
	}
	fmt.Println()
}

func runCmdMapB() {
	if len(mc.prevUrl) == 0 {
		fmt.Println("you are on the first page")
	} else {
		areas := pokeapi.GetLocationAreas(mc.prevUrl)
		mc.updateMapConfig(areas.Next, areas.Previous)
		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}
	}
	fmt.Println()
}

func runCmdExplore(location string) {
	if len(location) == 0 {
		fmt.Print("please enter area\n\n")
		return
	}
	url := fmt.Sprintf("%v%v", locationAreaUrl, location)
	area, err := pokeapi.GetLocationArea(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("exploring %v...\n", area.Name)
		fmt.Println("found pokemon:")
		for _, encounter := range area.PokemonEncounters {
			fmt.Printf("- %v\n", encounter.Pokemon.Name)
		}
	}
	fmt.Println()
}

const pokemonUrl = "https://pokeapi.co/api/v2/pokemon/"

var pokedex = map[string]pokeapi.Pokemon{}

func runCmdCatch(pokemon string) {
	if len(pokemon) == 0 {
		fmt.Print("please enter pokemon name\n\n")
		return
	}
	url := fmt.Sprintf("%v%v", pokemonUrl, pokemon)
	p, err := pokeapi.GetPokemon(url)
	if err != nil {
		fmt.Printf("%v\n\n", err)
		return
	}
	if _, ok := pokedex[p.Name]; ok {
		fmt.Printf("%v already in your pokedex!\n\n", p.Name)
		return
	}
	fmt.Printf("throwing a pokeball at %v...\n", p.Name)
	if rand.Intn(2) == 1 {
		pokedex[p.Name] = p
		fmt.Printf("%v was caught!\n", p.Name)
		fmt.Println("you may now inspect it with the 'inspect' command.")
	} else {
		fmt.Printf("%v escaped!\n", p.Name)
	}
	fmt.Println()
}

func runCmdInspect(pokemon string) {
	if len(pokemon) == 0 {
		fmt.Print("please enter pokemon name\n\n")
		return
	}
	p, ok := pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("name:", p.Name)
		fmt.Println("height:", p.Height)
		fmt.Println("weight:", p.Weight)
		fmt.Println("stats:")
		for _, s := range p.Stats {
			fmt.Printf("- %v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("types:")
		for _, t := range p.Types {
			fmt.Printf("- %v\n", t.Type.Name)
		}
	}
	fmt.Println()
}

func runCmdPokedex() {
	if len(pokedex) == 0 {
		fmt.Println("you have not caught any pokemon")
	} else {
		fmt.Println("your pokedex:")
		for _, p := range pokedex {
			fmt.Printf("- %v\n", p.Name)
		}
	}
	fmt.Println()
}
