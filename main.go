package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	}
	sc := bufio.NewScanner(os.Stdin)
	conf := &config{
		Previous: "",
		Next:     "",
	}
	for {
		fmt.Print("Pokedex > ")
		if !sc.Scan() {
			break
		}
		cmd := sc.Text()
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
		}
	}
}
