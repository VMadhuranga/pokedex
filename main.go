package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
	"pokedex/internal/pokecli"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func printPrompt() {
	fmt.Print("Pokedex > ")
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func main() {
	cliConfig := pokecli.CliConfig{
		RootLocationAreaUrl: "https://pokeapi.co/api/v2/location-area/",
		RootPokemonUrl:      "https://pokeapi.co/api/v2/pokemon/",
		Cache:               pokecache.NewCache(10 * time.Second),
		Pokedex:             map[string]pokeapi.Pokemon{},
		CliCommands: []pokecli.CliCommand{
			{
				Name:        "exit",
				Description: "Exit the Pokedex",
			},
			{
				Name:        "help",
				Description: "Displays a help message",
			},
			{
				Name:        "map",
				Description: "Displays next location areas",
			},
			{
				Name:        "mapb",
				Description: "Displays previous location areas",
			},
			{
				Name:        "explore",
				Description: "Displays list of all the Pokémon in location area",
			},
			{
				Name:        "catch",
				Description: "Catch Pokemon and adds them to the user's Pokedex",
			},
			{
				Name:        "inspect",
				Description: "Dispays details about a Pokemon",
			},
			{
				Name:        "pokedex",
				Description: "Dispays all the names of the Pokemon the user has caught ",
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	printPrompt()
	for scanner.Scan() {
		input := cleanInput(scanner.Text())

		switch input[0] {
		case "exit":
			err := cliConfig.InitExitCmd()
			if err != nil {
				handleError(err)
			}
		case "help":
			err := cliConfig.InitHelpCmd()
			if err != nil {
				handleError(err)
			}
		case "map":
			err := cliConfig.InitMapCmd()
			if err != nil {
				handleError(err)
			}
		case "mapb":
			err := cliConfig.InitMapbCmd()
			if err != nil {
				handleError(err)
			}
		case "explore":
			err := cliConfig.InitExploreCmd(input[1])
			if err != nil {
				handleError(err)
			}
		case "catch":
			err := cliConfig.InitCatchCmd(input[1])
			if err != nil {
				handleError(err)
			}
		case "inspect":
			err := cliConfig.InitInspectCmd(input[1])
			if err != nil {
				handleError(err)
			}
		case "pokedex":
			err := cliConfig.InitPokedexCmd()
			if err != nil {
				handleError(err)
			}
		default:
			fmt.Println("Unknown command")
		}

		printPrompt()
	}

	if err := scanner.Err(); err != nil {
		handleError(err)
	}
}
