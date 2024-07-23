package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/VMadhuranga/pokedex/pokecli"
)

func main() {
	pokeCliMap := pokecli.InitPokeCliMap()
	scanner := bufio.NewScanner(os.Stdin)
	pokecli.Prompt()
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		cmd, ok := pokeCliMap[input[0]]
		if !ok {
			err := errors.New("command not found, run 'help' for usage")
			fmt.Print(err, "\n\n")
		} else {
			arg := ""
			if len(input) > 1 {
				arg = input[1]
			}
			pokecli.RunCommand(pokeCliMap, cmd, arg)
		}
		pokecli.Prompt()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
