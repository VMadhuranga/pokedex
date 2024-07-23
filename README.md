# Pokedex

Pokedex is command-line REPL built with Go using [PokéAPI](https://pokeapi.co/) that you can catch and look up information about Pokemon

## Run Locally

- [Fork and clone](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo) the project
- Go to the project directory

  ```bash
  cd pokedex/
  ```

- Compile and run the program

  ```bash
  go build && ./pokedex
  ```

## Usage/Examples

```bash
pokedex > <command> [argument]
```

Command reference:

- `catch [pokemon]` : catch Pokemon and add them to the pokedex
- `exit` : exit the pokedex
- `explore [area]` : displays list of all the pokemon in a given area
- `help` : displays a help message
- `inspect [pokemon]` : displays details about pokemon
- `map` : displays the names of next 20 location areas
- `mapb` : displays the names of previous 20 location areas
- `pokedex` : displays list of captured pokemon

## Acknowledgements

This project is a part of [BOOT.DEV](https://www.boot.dev/), an online course to learn back-end development.
