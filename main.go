package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Cheemx/pokedexcli/internal/pokeapi"
	"github.com/Cheemx/pokedexcli/internal/pokemon"
)

var catchedPokemons = make(map[string]pokemon.Pokemon)

func main() {
	mep := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location in Pokemon World",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays names of pokemons in given location areas",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon specified by command argument",
			callback:    commandCatch,
		},
	}
	sc := bufio.NewScanner(os.Stdin)
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := &config{
		Previous: "",
		Next:     "",
		Client:   pokeClient,
	}
	for {
		fmt.Print("Pokedex > ")
		if !sc.Scan() {
			break
		}
		cmnd := sc.Text()
		cmdArgs := cleanInput(cmnd)
		if len(cmdArgs) < 1 {
			fmt.Println("Enter a command!")
		}
		cmd := cmdArgs[0]
		if (cmd == "explore" || cmd == "catch") && len(cmdArgs) == 1 {
			fmt.Println("You need to provide a valid location")
			continue
		}
		var argName string
		if len(cmdArgs) > 1 {
			argName = cmdArgs[1]
		}
		val, ok := mep[cmd]
		if !ok {
			fmt.Println("Unknown command")
		}

		switch val.name {
		case "help":
			val.callback(&config{})
			for k, v := range mep {
				fmt.Printf("%s: %s\n", k, v.description)
			}
		case "exit":
			val.callback(&config{})
		case "map":
			val.callback(conf)
		case "mapb":
			val.callback(conf)
		case "explore":
			val.callback(&config{
				Previous: "",
				Next:     argName,
				Client:   pokeClient,
			})
		case "catch":
			val.callback(&config{
				Previous: "",
				Next:     argName,
				Client:   pokeClient,
			})
		}
	}
}
